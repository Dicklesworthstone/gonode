// Code generated by goa v3.3.1, DO NOT EDIT.
//
// walletnode HTTP client CLI support package
//
// Command:
// $ goa gen github.com/pastelnetwork/walletnode/api/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	artworksc "github.com/pastelnetwork/walletnode/api/gen/http/artworks/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `artworks (register|register-task-state|register-task|register-tasks|upload-image)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` artworks register --body '{
      "artist_name": "Leonardo da Vinci",
      "artist_pastelid": "jXYJud3rmrR1Sk2scvR47N4E4J5Vv48uCC6se2nzHrBRdjaKj3ybPoi1Y2VVoRqi1GnQrYKjSxQAC7NBtvtEdS",
      "artist_pastelid_passphrase": "qwerasdf1234",
      "artist_website_url": "https://www.leonardodavinci.net",
      "description": "The Mona Lisa is an oil painting by Italian artist, inventor, and writer Leonardo da Vinci. Likely completed in 1506, the piece features a portrait of a seated woman set against an imaginary landscape.",
      "image_id": 1,
      "issued_copies": 1,
      "keywords": "Renaissance, sfumato, portrait",
      "maximum_fee": 100,
      "name": "Mona Lisa",
      "series_name": "Famous artist",
      "spendable_address": "PtiqRXn2VQwBjp1K8QXR2uW2w2oZ3Ns7N6j",
      "youtube_url": "https://www.youtube.com/watch?v=0xl6Ufo4ZX0"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
	dialer goahttp.Dialer,
	artworksConfigurer *artworksc.ConnConfigurer,
	artworksUploadImageEncoderFn artworksc.ArtworksUploadImageEncoderFunc,
) (goa.Endpoint, interface{}, error) {
	var (
		artworksFlags = flag.NewFlagSet("artworks", flag.ContinueOnError)

		artworksRegisterFlags    = flag.NewFlagSet("register", flag.ExitOnError)
		artworksRegisterBodyFlag = artworksRegisterFlags.String("body", "REQUIRED", "")

		artworksRegisterTaskStateFlags      = flag.NewFlagSet("register-task-state", flag.ExitOnError)
		artworksRegisterTaskStateTaskIDFlag = artworksRegisterTaskStateFlags.String("task-id", "REQUIRED", "Task ID of the registration process")

		artworksRegisterTaskFlags      = flag.NewFlagSet("register-task", flag.ExitOnError)
		artworksRegisterTaskTaskIDFlag = artworksRegisterTaskFlags.String("task-id", "REQUIRED", "Task ID of the registration process")

		artworksRegisterTasksFlags = flag.NewFlagSet("register-tasks", flag.ExitOnError)

		artworksUploadImageFlags    = flag.NewFlagSet("upload-image", flag.ExitOnError)
		artworksUploadImageBodyFlag = artworksUploadImageFlags.String("body", "REQUIRED", "")
	)
	artworksFlags.Usage = artworksUsage
	artworksRegisterFlags.Usage = artworksRegisterUsage
	artworksRegisterTaskStateFlags.Usage = artworksRegisterTaskStateUsage
	artworksRegisterTaskFlags.Usage = artworksRegisterTaskUsage
	artworksRegisterTasksFlags.Usage = artworksRegisterTasksUsage
	artworksUploadImageFlags.Usage = artworksUploadImageUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "artworks":
			svcf = artworksFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "artworks":
			switch epn {
			case "register":
				epf = artworksRegisterFlags

			case "register-task-state":
				epf = artworksRegisterTaskStateFlags

			case "register-task":
				epf = artworksRegisterTaskFlags

			case "register-tasks":
				epf = artworksRegisterTasksFlags

			case "upload-image":
				epf = artworksUploadImageFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "artworks":
			c := artworksc.NewClient(scheme, host, doer, enc, dec, restore, dialer, artworksConfigurer)
			switch epn {
			case "register":
				endpoint = c.Register()
				data, err = artworksc.BuildRegisterPayload(*artworksRegisterBodyFlag)
			case "register-task-state":
				endpoint = c.RegisterTaskState()
				data, err = artworksc.BuildRegisterTaskStatePayload(*artworksRegisterTaskStateTaskIDFlag)
			case "register-task":
				endpoint = c.RegisterTask()
				data, err = artworksc.BuildRegisterTaskPayload(*artworksRegisterTaskTaskIDFlag)
			case "register-tasks":
				endpoint = c.RegisterTasks()
				data = nil
			case "upload-image":
				endpoint = c.UploadImage(artworksUploadImageEncoderFn)
				data, err = artworksc.BuildUploadImagePayload(*artworksUploadImageBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// artworksUsage displays the usage of the artworks command and its subcommands.
func artworksUsage() {
	fmt.Fprintf(os.Stderr, `Pastel Artwork
Usage:
    %s [globalflags] artworks COMMAND [flags]

COMMAND:
    register: Runs a new registration process for the new artwork.
    register-task-state: Streams the state of the registration process.
    register-task: Returns a single task.
    register-tasks: List of all tasks.
    upload-image: Upload the image that is used when registering a new artwork.

Additional help:
    %s artworks COMMAND --help
`, os.Args[0], os.Args[0])
}
func artworksRegisterUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] artworks register -body JSON

Runs a new registration process for the new artwork.
    -body JSON: 

Example:
    `+os.Args[0]+` artworks register --body '{
      "artist_name": "Leonardo da Vinci",
      "artist_pastelid": "jXYJud3rmrR1Sk2scvR47N4E4J5Vv48uCC6se2nzHrBRdjaKj3ybPoi1Y2VVoRqi1GnQrYKjSxQAC7NBtvtEdS",
      "artist_pastelid_passphrase": "qwerasdf1234",
      "artist_website_url": "https://www.leonardodavinci.net",
      "description": "The Mona Lisa is an oil painting by Italian artist, inventor, and writer Leonardo da Vinci. Likely completed in 1506, the piece features a portrait of a seated woman set against an imaginary landscape.",
      "image_id": 1,
      "issued_copies": 1,
      "keywords": "Renaissance, sfumato, portrait",
      "maximum_fee": 100,
      "name": "Mona Lisa",
      "series_name": "Famous artist",
      "spendable_address": "PtiqRXn2VQwBjp1K8QXR2uW2w2oZ3Ns7N6j",
      "youtube_url": "https://www.youtube.com/watch?v=0xl6Ufo4ZX0"
   }'
`, os.Args[0])
}

func artworksRegisterTaskStateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] artworks register-task-state -task-id INT

Streams the state of the registration process.
    -task-id INT: Task ID of the registration process

Example:
    `+os.Args[0]+` artworks register-task-state --task-id 5
`, os.Args[0])
}

func artworksRegisterTaskUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] artworks register-task -task-id INT

Returns a single task.
    -task-id INT: Task ID of the registration process

Example:
    `+os.Args[0]+` artworks register-task --task-id 5
`, os.Args[0])
}

func artworksRegisterTasksUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] artworks register-tasks

List of all tasks.

Example:
    `+os.Args[0]+` artworks register-tasks
`, os.Args[0])
}

func artworksUploadImageUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] artworks upload-image -body JSON

Upload the image that is used when registering a new artwork.
    -body JSON: 

Example:
    `+os.Args[0]+` artworks upload-image --body '{
      "file": "VmVsIHZvbHVwdGF0ZW0gcHJvdmlkZW50IGRvbG9yaWJ1cy4="
   }'
`, os.Args[0])
}
