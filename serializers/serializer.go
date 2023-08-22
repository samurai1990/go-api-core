package serializers

type GetSerializeResponseInterface interface {
	ListSerializerResponse(model any) (any, error)
}
type SigninSerializerResponseInterface interface {
	SigninSerializerResponse(model any) error
}

type CreateSerializerResponseInterface interface {
	CreateSerializerResponse(model any) error
}

type RetrieveSerializerResponseInterface interface {
	RetrieveSerializerResponse(model any) error
}
type UpdateSerializerResponseInterface interface {
	UpdateSerializerResponse(model any) error
}

type Serializer struct {
	model   any
	SerData any
}

func NewSerializer(model any) *Serializer {
	return &Serializer{
		model: model,
	}
}

func (s *Serializer) GetListSerializer(i GetSerializeResponseInterface) error {
	resser, err := i.ListSerializerResponse(s.model)
	if err != nil {
		return err
	}
	s.SerData = resser
	return nil
}

func (s *Serializer) GetRetrieveSerializer(i RetrieveSerializerResponseInterface) error {
	if err := i.RetrieveSerializerResponse(s.model); err != nil {
		return err
	} else {
		return nil
	}
}

func (s *Serializer) GetUpdateSerializer(i UpdateSerializerResponseInterface) error {
	if err := i.UpdateSerializerResponse(s.model); err != nil {
		return err
	} else {
		return nil
	}
}

func (s *Serializer) GetSigninSerializer(i SigninSerializerResponseInterface) error {
	err := i.SigninSerializerResponse(s.model)
	if err != nil {
		return err
	}
	return nil
}

func (s *Serializer) GetCreateSerializer(i CreateSerializerResponseInterface) error {
	err := i.CreateSerializerResponse(s.model)
	if err != nil {
		return err
	}
	return nil
}
