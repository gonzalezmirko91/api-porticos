package env

type Environment string

const (
	Development Environment = "development"
	Dev         Environment = "dev"
	Production  Environment = "production"
	Prod        Environment = "prod"
	Qa          Environment = "qa"
	UAT         Environment = "UAT"
)

func (e Environment) IsDev() bool {
	return e == Development || e == Dev
}

func (e Environment) IsProduction() bool {
	return e == Prod || e == Production
}

func (e Environment) IsQA() bool {
	return e == Qa || e == UAT
}

func Parse(env string) Environment {
	switch env {
	case "development", "dev":
		return Development
	case "production", "prod":
		return Production
	case "qa", "UAT":
		return Qa
	default:
		return Development // Fallback seguro
	}
}
