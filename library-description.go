package main 


import (
	"fmt"
)

type LibraryDescriptionOptions struct {
	Name      string   `short:"n" long:"name" description:"Name of the library"`
	Version   string   `short:"v" long "version" description:"Version of the library"`
	Includes  []string `short:"I" long:"include" description:"Path to append to the list of includes"`
	Libdirs   []string `short:"L" description:"Path for finding libraries"`
	Libs      []string `short:"l" description:"Library to add"`
	PkgConfig string   `short:"P" long:"pkg-config" description:"Pkg config for extracting name, version, includes and libs"`
	Deb       []string `short:"D" long:"deb" description:"Debian package to extract al data from, all path will be expanded from"`
}



func (ld *LibraryDescriptionOptions) Format(f fmt.State, c rune) {
	fmt.Fprintf(f,"Name: %s, Version: %s, Includes: %s Libdirs: %s Libs:%s PkgConfig:%s, Debs: %s",
		ld.Name,
		ld.Version,
		ld.Includes,
		ld.Libdirs, 
		ld.Libs, 
		ld.PkgConfig,
		ld.Deb)
}




func (ld *LibraryDescriptionOptions) IsValid() (bool,error) {
	if len(ld.PkgConfig) != 0 {
		return true,nil
	}
	var reasons []string
	reasons = nil
	if len(ld.Name) == 0 {
		reasons = append(reasons,"You need to provide a name or let pkg-config define it for you")
	}
	if len(ld.Version) == 0 {
		reasons = append(reasons,"You need to provide a version or let pkg-config define it for you")
	}

	if len(ld.Includes) == 0 {
		reasons = append(reasons, "You need to provide at least one Include dir")
	}

	if len(ld.Libs) == 0 {
		reasons = append(reasons, "You need to provide at least one Library")
	}

	if len(reasons) == 0 {
		return true, nil
	}
	return false,fmt.Errorf("Invalid definition : %s",reasons)
}

func (ld *LibraryDescriptionOptions) Execute(args []string) error {
	if _,err := opts.IsValid(); err != nil {
		return err
	}
	if _,err := ld.IsValid(); err != nil {
		return err
	}
	fmt.Println("I compare this new library :")
	fmt.Println(&opts)
	fmt.Println("To this new one :")
	fmt.Println(ld)
	return nil
}


func init(){
	parser.AddCommand("compare-to",
		"Compare the library to another one",
		"Compare the library defined by the library before the command to the one after the command. Definition are transfered",
		&LibraryDescriptionOptions{})

}
