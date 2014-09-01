package migrate

type rev20140828205400 struct{}

var ContainersTable = &rev20140828205400{}

func (r *rev20140828205400) Revision() int64 {
	return 20140828205400
}

func (r *rev20140828205400) Up(mg *MigrationDriver) error {
    _, err := mg.CreateTable("containers", []string{
        mg.T.Integer("build_id", PRIMARYKEY),
        mg.T.Blob("containers"),
    });
    return err
}

func (r *rev20140828205400) Down(mg *MigrationDriver) error {
    _, err := mg.DropTable("containers");
    return err
}
