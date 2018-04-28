package encoder

type Statement struct {
	Sid string `hcl:"sid"`
	Effect string `hcl:"effect"`
}

type DataSource struct {
	Type string `hcl:",key"`
	Name string `hcl:",key"`
	Statements []Statement `hcl:"statement,squash"`
}

type WholeFile struct {
	Entry DataSource `hcl:"data"`
}
