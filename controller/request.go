package handler

import "fmt"

func errParamIsrequired(name string, typ string) error {
	return fmt.Errorf("param: %s (%s) is required", name, typ)
}

type CreatePersonRequest struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Ocupation string `json:"ocupation"`
	Salary    int    `json:"salary"`
	Alive     *bool  `json:"alive"`
}

func (r *CreatePersonRequest) Validate() error {
	if r.Name == "" && r.Gender == "" && r.Ocupation == "" && r.Alive == nil && r.Age <= 0 && r.Salary <= 0{
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return errParamIsrequired("name", "string")
	}
	if r.Gender == "" {
		return errParamIsrequired("gender", "string")
	}
	if r.Ocupation == "" {
		return errParamIsrequired("ocupation", "string")
	}
	if r.Alive == nil {
		return errParamIsrequired("alive", "boolean")
	}
	if r.Salary < 0 {
		return errParamIsrequired("salary", "int")
	}
	if r.Age < 0 {
		return errParamIsrequired("age", "int")
	}
	return nil
}
