package migrate

type rev20140828205400 struct{}

var ContainersTable = &rev20140828205400{}

func (r *rev20140828205400) Revision() int64 {
	return 20140828205400
}

func (r *rev20140828205400) Up(mg *MigrationDriver) error {
    if _, err := mg.CreateTable("containers", []string{
        t.Integer("build_id", PRIMARYKEY),
        t.Blob("containers"),
    }); err != nil {
        return err
    }

}

func (r *rev20140828205400) Down(mg *MigrationDriver) error {
    if _, err := mg.DropTable("containers"); err != nil {
        return err
    }
}
