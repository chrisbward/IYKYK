# IYKYK

Check the test cases, and you'll have an idea what this is useful for.

## License

This code is released under the Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International license.
Get in touch if you wish to license this code for commercial purposes.

## Usage

```go

import github.com/chrisbward/IYKYK

func stripAction() {

    // contains the utility methods for stripping the content
    stripContentController, _ := stripcontentcontroller.NewStripContentController()

    contentController, _ := contentcontroller.NewContentController(ContentControllerOptions{
        StripEmDash:       true,
        StripEmoji:        true,
        StripAngledQuotes: true,
    }, stripContentController) 

    input := `Here is some content — I wish to be “cleaned”. It’s very useful 🚀 for certain purposes `

    output, err := contentController.CleanContentAutomatic(input)
    fmt.Println(output)
}

```
Expected result;

```
Here is some content - I wish to be "cleaned". It's very useful for certain purposes
```


## Tests

To run the unit tests;
```bash
make test
```
