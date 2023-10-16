package models

type Provider struct {
	ProviderId int `json:"provider_id"`

	ProviderName string `json:"provider_name"`

	INN string `json:"INN"`

	ContactDetails string `json:"contact_details"`

	RF string `json:"RF"`
}

type Product struct {
	ProviderId int `json:"provider_id"`

	ElementId int `json:"element_id"`

	ElementName string `json:"element_name"`

	Quantity string `json:"quantity"`

	Price string `json:"price"`

	Amount string `json:"amount"`

	Deadline string `json:"deadline"`
}

/*
*
In GO exported variable and function name
must start with big cap
*/
func SampleProducts() []Provider {

	var products = []Provider{

		{ProviderId: 8, ProviderName: "P0122", INN: "123456789987", ContactDetails: "122", RF: "yes"},
		{ProviderId: 9, ProviderName: "P0123", INN: "123456789986", ContactDetails: "69", RF: "yes"},
		{ProviderId: 10, ProviderName: "P0124", INN: "123456789985", ContactDetails: "280", RF: "yes"},
	}

	return products
}
