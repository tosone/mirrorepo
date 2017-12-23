package config

type config struct {
	DBEngine  string `yaml:"databaseEngine"`
	DBPath    string `yaml:"databasePath"`
	Log       string `yaml:"log"`
	MaxThread uint   `yaml:"maxThread"`
	Repo      string `yaml:"repo"`
}
