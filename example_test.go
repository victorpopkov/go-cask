package cask

import "fmt"

func Example_one() {
	// for this example we will load the cask from our testdata directory
	content := string(getTestdata("example-one.rb"))

	// example
	c := NewCask(content)
	err := c.Parse()

	if err == nil {
		fmt.Println("Token:", c.Token)
		for i, v := range c.Variants {
			fmt.Printf("Variant #%d:\n", i+1)
			fmt.Printf("%10s: %s\n", "version", v.Version)
			fmt.Printf("%10s: %s\n", "sha256", v.SHA256)
			fmt.Printf("%10s: %s\n", "url", v.URL)
			fmt.Printf("%10s: %s\n", "appcast", v.Appcast.URL)
			fmt.Printf("%12s%s\n", "", v.Appcast.Checkpoint)
			fmt.Printf("%10s: %v\n", "names", v.Names)
			fmt.Printf("%10s: %s\n", "homepage", v.Homepage)

			// artifacts
			fmt.Printf("%10s: ", "artifacts")
			if len(v.Artifacts) > 0 {
				for i, a := range v.Artifacts {
					if i == 0 {
						fmt.Printf("%s\n", a.String())
					} else {
						fmt.Printf("%12s%s\n", "", a.String())
					}
				}
			} else {
				fmt.Printf("%v\n", v.Artifacts)
			}

			// macOS
			fmt.Printf("%10s: %s [minimum]\n", "macOS", v.MinimumSupportedMacOS)
			fmt.Printf("%12s%s [maximum]\n", "", v.MaximumSupportedMacOS)
		}
	}

	// Output:
	// Token: example-one
	// Variant #1:
	//    version: 2.0.0
	//     sha256: f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261
	//        url: https://example.com/app_#{version}.dmg
	//    appcast: https://example.com/sparkle/#{version.major}/appcast.xml
	//             8dc47a4bcec6e46b79fb6fc7b84224f1461f18a2d9f2e5adc94612bb9b97072d
	//      names: [Example Example One]
	//   homepage: https://example.com/
	//  artifacts: app, Example One.app => Example.app
	//             app, Example One Uninstaller.app
	//             binary, #{appdir}/Example One.app/Contents/MacOS/example-one => example
	//      macOS: macOS High Sierra (10.13) [minimum]
	//             macOS High Sierra (10.13) [maximum]
}

func Example_two() {
	// for this example we will load the cask from our testdata directory
	content := string(getTestdata("example-two.rb"))

	// example
	c := NewCask(content)
	err := c.Parse()

	if err == nil {
		fmt.Println("Token:", c.Token)
		for i, v := range c.Variants {
			fmt.Printf("Variant #%d:\n", i+1)
			fmt.Printf("%10s: %s\n", "version", v.Version)
			fmt.Printf("%10s: %s\n", "sha256", v.SHA256)
			fmt.Printf("%10s: %s\n", "url", v.URL)
			fmt.Printf("%10s: %s\n", "appcast", v.Appcast.URL)
			fmt.Printf("%12s%s\n", "", v.Appcast.Checkpoint)
			fmt.Printf("%10s: %v\n", "names", v.Names)
			// fmt.Printf("%10s: %s\n", "homepage", v.Homepage)

			// artifacts
			fmt.Printf("%10s: ", "artifacts")
			if len(v.Artifacts) > 0 {
				for i, a := range v.Artifacts {
					if i == 0 {
						fmt.Printf("%s\n", a.String())
					} else {
						fmt.Printf("%12s%s\n", "", a.String())
					}
				}
			} else {
				fmt.Printf("%v\n", v.Artifacts)
			}

			// macOS
			fmt.Printf("%10s: %s [minimum]\n", "macOS", v.MinimumSupportedMacOS)
			fmt.Printf("%12s%s [maximum]\n", "", v.MaximumSupportedMacOS)
		}
	}

	// Output:
	// Token: example-two
	// Variant #1:
	//    version: 1.5.0
	//     sha256: 1f4dc096d58f7d21e3875671aee6f29b120ab84218fa47db2cb53bc9eb5b4dac
	//        url: https://example.com/app_#{version}.pkg
	//    appcast: https://example.com/sparkle/#{version}/el_capitan.xml
	//             93ef3101ca730028d70524f71b7f6f17cbdb8d26906299f90c38b7079e1d03a4
	//      names: []
	//  artifacts: []
	//      macOS: Mac OS X Tiger (10.4) [minimum]
	//             OS X El Capitan (10.11) [maximum]
	// Variant #2:
	//    version: 2.0.0
	//     sha256: f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261
	//        url: https://example.com/app_#{version}.pkg
	//    appcast: https://example.com/sparkle/#{version.major}/appcast.xml
	//             57956bd3fb23a5673e30dc83ed19d51b43e5a9235756887f3ed90662e6c68fb7
	//      names: [Example Example Two]
	//  artifacts: pkg, app_#{version}.pkg, allow_untrusted: true
	//      macOS: macOS High Sierra (10.13) [minimum]
	//             macOS High Sierra (10.13) [maximum]
}
