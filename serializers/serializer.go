package serializers

type GetSerializeResponseInterface interface {
	ListSerializerResponse(model any) (any, error)
}
type SigninSerializerResponseInterface interface {
	SigninSerializerResponse(model any) error
}

type GetSerializer struct {
	model any
	SerData any
}

func NewGetSerializer(model any) *GetSerializer {
	return &GetSerializer{
		model: model,
	}
}

func (s *GetSerializer) GetListSerializer(i GetSerializeResponseInterface) error {
	resser,err := i.ListSerializerResponse(s.model)
	if err != nil {
		return err
	}
	s.SerData=resser
	return nil
}
func (s *GetSerializer) GetSigninSerializer(i SigninSerializerResponseInterface) error {
	err := i.SigninSerializerResponse(s.model)
	if err != nil {
		return err
	}
	return nil
}
