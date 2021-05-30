package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/gitchander/permutation"
	"github.com/maxhora/goptuna"
	"golang.org/x/sync/errgroup"
	"gorm.io/driver/mysql"

	"github.com/maxhora/goptuna/rdb.v2"
	"github.com/maxhora/goptuna/tpe"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/pastelnetwork/gonode/common/errors"
	"github.com/pastelnetwork/gonode/dupe-detection/pkg/auprc"
	"github.com/pastelnetwork/gonode/dupe-detection/pkg/dupedetection"
)

const cacheFileName = "cached"

var evaluateNumberOfTimes = 500
var rootDir = ""
var numberOfImagesToValidate = 0

// objective defines the objective of the study - find out the best aurpc value
func objective(trial goptuna.Trial) (float64, error) {
	fmt.Println("Start objective...")
	var err error
	// Define the search space via Suggest APIs.
	config := dupedetection.NewComputeConfig()

	config.RootDir = rootDir
	err = trial.SetUserAttr("RootDir", config.RootDir)
	if err != nil {
		return 0, errors.New(err)
	}

	g, _ := errgroup.WithContext(context.Background())

	err = trial.SetUserAttr("TrimByPercentile", fmt.Sprintf("%v", config.TrimByPercentile))
	if err != nil {
		return 0, errors.New(err)
	}

	config.NumberOfImagesToValidate = numberOfImagesToValidate
	err = trial.SetUserAttr("MaxImageCountToEvaluate", fmt.Sprintf("%v", config.NumberOfImagesToValidate*2))
	if err != nil {
		return 0, errors.New(err)
	}

	g.Go(func() error {
		fmt.Println("Suggesting MIThreshold...")
		config.MIThreshold, err = trial.SuggestFloat("MIThreshold", 5.2, 5.4)
		if err != nil {
			return errors.New(err)
		}
		return nil
	})

	g.Go(func() error {
		fmt.Println("Suggesting Pearson...")
		config.PearsonDupeThreshold, err = trial.SuggestFloat("Pearson", 0.99, 0.99999)
		if err != nil {
			return errors.New(err)
		}
		return nil
	})

	g.Go(func() error {
		fmt.Println("Suggesting Spearman...")
		config.SpearmanDupeThreshold, err = trial.SuggestFloat("Spearman", 0.75, 0.85)
		if err != nil {
			return errors.New(err)
		}
		return nil
	})

	g.Go(func() error {
		fmt.Println("Suggesting Kendall...")
		config.KendallDupeThreshold, _ = trial.SuggestFloat("Kendall", 0.68, 0.72)
		if err != nil {
			return errors.New(err)
		}
		return nil
	})
	/*config.RandomizedDependenceDupeThreshold, _ = trial.SuggestFloat("RDC", 0.5, 0.99999)
	if err != nil {
		return 0, errors.New(err)
	}*/

	g.Go(func() error {
		fmt.Println("Suggesting Hoeffding...")
		config.HoeffdingDupeThreshold, _ = trial.SuggestFloat("Hoeffding", 0.2, 0.6)
		if err != nil {
			return errors.New(err)
		}
		return nil
	})

	g.Go(func() error {
		fmt.Println("Suggesting Blomqvist...")
		config.BlomqvistDupeThreshold, _ = trial.SuggestFloat("Blomqvist", 0.6, 0.8)
		if err != nil {
			return errors.New(err)
		}
		return nil
	})
	/*config.HoeffdingDupeThreshold, _ = trial.SuggestFloat("HoeffdingD1", 0.1, 0.99999)
	if err != nil {
		return 0, errors.New(err)
	}
	config.HoeffdingRound2DupeThreshold, _ = trial.SuggestFloat("HoeffdingD2", 0.1, 0.99999)
	if err != nil {
		return 0, errors.New(err)
	}*/

	/*allCombinationsOfUnstableMethods := combinations.All(config.UnstableOrderOfCorrelationMethods)
	var allOrderedCombinationsOfUnstableMethodsAsStrings []string
	for _, combination := range allCombinationsOfUnstableMethods {
		permutator := permutation.New(permutation.StringSlice(combination))
		for permutator.Next() {
			allOrderedCombinationsOfUnstableMethodsAsStrings = append(allOrderedCombinationsOfUnstableMethodsAsStrings, strings.Join(combination, " "))
		}
	}*/

	var allOrderedCombinationsOfMethodsAsStrings []string
	combination := config.CorrelationMethodNameArray
	permutator := permutation.New(permutation.StringSlice(combination))
	for permutator.Next() {
		allOrderedCombinationsOfMethodsAsStrings = append(allOrderedCombinationsOfMethodsAsStrings, strings.Join(combination, " "))
	}

	g.Go(func() error {
		fmt.Println("Suggesting CorrelationMethodsOrderIndex...")
		correlationMethodIndex, err := trial.SuggestDiscreteFloat("CorrelationMethodsOrderIndex", 0, float64(len(allOrderedCombinationsOfMethodsAsStrings)-1), 1.0)
		config.CorrelationMethodsOrder = allOrderedCombinationsOfMethodsAsStrings[int(correlationMethodIndex)]
		if err != nil {
			return errors.New(err)
		}
		return nil
	})
	//correlationMethodsOrder := append(config.StableOrderOfCorrelationMethods, allOrderedCombinationsOfMethodsAsStrings[int(correlationMethodIndex)])
	//config.CorrelationMethodsOrder = strings.Join(correlationMethodsOrder, " ")

	//config.CorrelationMethodsOrder = "MI PearsonR SpearmanRho BootstrappedKendallTau BootstrappedBlomqvistBeta HoeffdingDRound1 HoeffdingDRound2"
	//config.CorrelationMethodsOrder = "PearsonR SpearmanRho KendallTau HoeffdingD BlomqvistBeta"

	err = trial.SetUserAttr("CorrelationMethodsOrder", config.CorrelationMethodsOrder)
	if err != nil {
		return 0, errors.New(err)
	}

	err = g.Wait()
	if err != nil {
		return 0, errors.New(err)
	}

	fmt.Println("Suggesting ended. Measure AUPRC...")

	aurpcResult, err := auprc.MeasureAUPRC(config)
	if err != nil {
		return 0, errors.New(err)
	}
	err = trial.SetUserAttr("DupeAccuracy", fmt.Sprintf("%v", aurpcResult.DupeAccuracy))
	if err != nil {
		return 0, errors.New(err)
	}
	err = trial.SetUserAttr("DupeCount", fmt.Sprintf("%v", aurpcResult.DupeCount))
	if err != nil {
		return 0, errors.New(err)
	}
	err = trial.SetUserAttr("OriginalAccuracy", fmt.Sprintf("%v", aurpcResult.OriginalAccuracy))
	if err != nil {
		return 0, errors.New(err)
	}
	err = trial.SetUserAttr("OriginalCount", fmt.Sprintf("%v", aurpcResult.OriginalCount))
	if err != nil {
		return 0, errors.New(err)
	}
	err = trial.SetUserAttr("AverageAccuracy", fmt.Sprintf("%v", aurpcResult.AverageAccuracy))
	if err != nil {
		return 0, errors.New(err)
	}
	return 1.0 - aurpcResult.AUPRC, nil
}

func runStudy(studyName string) error {
	db, err := gorm.Open(mysql.Open("goptuna:password@tcp(localhost:3306)/goptuna?parseTime=true"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return errors.New(err)
	}

	storage := rdb.NewStorage(db)

	// Creates new study or loads available
	study, err := goptuna.CreateStudy(
		studyName,
		goptuna.StudyOptionStorage(storage),
		//goptuna.StudyOptionRelativeSampler(cmaes.NewSampler(cmaes.SamplerOptionIPop(2))),
		goptuna.StudyOptionSampler(tpe.NewSampler()),
		goptuna.StudyOptionLoadIfExists(true),
	)
	if err != nil {
		return errors.New(err)
	}

	// Evaluate objective function specified number of times
	err = study.Optimize(objective, evaluateNumberOfTimes)
	if err != nil {
		return errors.New(err)
	}

	v, err := study.GetBestValue()
	if err != nil {
		return errors.New(err)
	}
	p, err := study.GetBestParams()
	if err != nil {
		return errors.New(err)
	}

	log.Printf("Best value=%f", v)
	for key, value := range p {
		log.Printf("%v=%v", key, value)
	}
	return nil
}

func main() {
	rootDirPtr := flag.String("rootDir", "", "a path to the directory with the test corpus of images.")
	goptunaStudyNamePtr := flag.String("studyName", "dupe-detection-aurpc", "a name of the Goptuna study to create or continue available.")
	numberOfImagesToValidatePtr := flag.Int("imageCount", 0, "limits the number of dupes and original images to validate.")
	evaluateNumberOfTimesPtr := flag.Int("runCount", 500, "defines the number of times goptuna will evaluate optimization objective.")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	rootDir = *rootDirPtr
	numberOfImagesToValidate = *numberOfImagesToValidatePtr
	evaluateNumberOfTimes = *evaluateNumberOfTimesPtr

	memoizer := dupedetection.GetMemoizer()
	memoizer.Storage.LoadFile(cacheFileName)

	if err := runStudy(*goptunaStudyNamePtr); err != nil {
		log.Print(errors.ErrorStack(err))
		panic(err)
	}

	memoizer.Storage.SaveFile(cacheFileName)
}
