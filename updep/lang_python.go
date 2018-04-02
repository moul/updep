package updep

var LangPython = Language{
	Name: "python",
}

func init() {
	Languages["python"] = LangPython
}
