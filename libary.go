package main


type Library struct {
	Name      string
	Version   string
	Includes []string
	Binaries []string
	explorer ArchiveExplorer
}



func NewLibrary(desc *LibraryDescriptionOptions) (*Library, error){
	l := Library{Name: desc.Name,Version: desc.Version,Includes: nil,Binaries: nil,explorer: nil}
	
	var err error

	if len(desc.Deb) > 0 {
		if l.explorer,err = NewDebExplorer(desc.Deb); err != nil {
			return nil,err
		}
	} else {
		l.explorer = &LocalFS{}
	}

	return &l,nil
}

func (l* Library) AddPkgConfig(p string) error {

	return nil;
}


func (l *Library) SetIncludes(includes []string) error {
	return nil
}
