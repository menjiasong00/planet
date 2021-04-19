package fuc

type BasFunc struct {}

func  (b *BasFunc) Contact(args ...string) string {
	s := ""
	for _,v := range args {
		s += v
	}
	return s
}

