package commands

func init() {

	New("timezone", "get").
		Returns(
			StringValue("timezone"),
		)
	New("timezone", "set").
		Accepts(
			StringArg("timezone").
				WithDocsDefault("Europe/Moscow").
				WithDescription("Timezone to set to in ISO 8601 format"),
		).
		Returns(
			StringValue("timezone"),
		)

	New("ntp", "get-list").NotImplemented()
	New("ntp", "add").NotImplemented().
		Accepts(StringArg("server")).
		NotImplemented()
	New("ntp", "remove").
		Accepts(StringArg("server")).
		NotImplemented()

	New("firewall", "get-status").NotImplemented()
	New("firewall", "enable").NotImplemented()
	New("firewall", "disable").NotImplemented()

	New("modems", "get-list").NotImplemented()
	New("modems", "enable").
		Accepts(StringArg("modem")).
		NotImplemented()
	New("modems", "disable").
		Accepts(StringArg("modem")).
		NotImplemented()
	New("modems", "get-signal").
		Accepts(StringArg("modem")).
		NotImplemented()

	New("system", "reboot").NotImplemented()

	New("ssh", "get-status").NotImplemented()
	New("ssh", "enable").NotImplemented()
	New("ssh", "disable").NotImplemented()

	New("journals", "get-journal").
		Accepts(
			StringArg("journal").
				WithValues("system", "core", "connection", "port-forwarding"),
			IntArg("lines").
				WithDocsDefault(100),
		).
		NotImplemented()

}
