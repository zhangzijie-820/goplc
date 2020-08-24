package tag

import (
	"bytes"
	"github.com/MiguelValentine/goplc/ethernetip/commonIndustrialProtocol"
	"github.com/MiguelValentine/goplc/ethernetip/commonIndustrialProtocol/segment/epath"
	"github.com/MiguelValentine/goplc/lib"
)

const ServiceReadTag = commonIndustrialProtocol.Service(0x4c)
const ServiceWriteTag = commonIndustrialProtocol.Service(0x4d)

type Tag struct {
	name      []byte
	readCount uint16
	xtype     DataType
	value     []byte
	OnChange  func(interface{})
	OnData    func(interface{})
	next      func()
}

func (t *Tag) GenerateReadMessageRequest() *commonIndustrialProtocol.MessageRouterRequest {
	mr := &commonIndustrialProtocol.MessageRouterRequest{}
	mr.Service = ServiceReadTag
	mr.RequestPath = epath.DataBuild(epath.DataTypeANSI, t.name, true)

	data := new(bytes.Buffer)
	lib.WriteByte(data, t.readCount)
	mr.RequestData = data.Bytes()

	return mr
}

func (t *Tag) GenerateWriteMessageRequest() *commonIndustrialProtocol.MessageRouterRequest {
	mr := &commonIndustrialProtocol.MessageRouterRequest{}
	mr.Service = ServiceWriteTag
	mr.RequestPath = epath.DataBuild(epath.DataTypeANSI, t.name, true)

	data := new(bytes.Buffer)
	lib.WriteByte(data, t.xtype)
	lib.WriteByte(data, t.readCount)
	lib.WriteByte(data, t.GetValue())
	mr.RequestData = data.Bytes()

	return mr
}

func (t *Tag) ReadTagParser(mr *commonIndustrialProtocol.MessageRouterResponse) {
	dataReader := bytes.NewReader(mr.ResponseData)
	lib.ReadByte(dataReader, &t.xtype)
	newValue := make([]byte, dataReader.Len())
	lib.ReadByte(dataReader, newValue)

	if t.OnChange != nil && bytes.Compare(t.value, newValue) != 0 {
		t.value = newValue
		t.OnChange(t.GetValue())
	}

	if t.OnData != nil {
		t.OnData(t.GetValue())
	}

	if t.next != nil {
		f := t.next
		t.next = nil
		f()
	}
}

func (t *Tag) WriteTagParser(mr *commonIndustrialProtocol.MessageRouterResponse) {
	if t.next != nil {
		f := t.next
		t.next = nil
		f()
	}
}

func (t *Tag) Then(f func()) {
	t.next = f
}

func (t *Tag) Type() string {
	return TypeMap[t.xtype]
}

func (t *Tag) Name() string {
	return string(t.name)
}

func NewTag(name string) *Tag {
	_tag := &Tag{}
	_tag.name = []byte(name)
	_tag.readCount = 1
	_tag.xtype = NULL
	return _tag
}

func NewTagWithType(name string, tp DataType) *Tag {
	_tag := &Tag{}
	_tag.name = []byte(name)
	_tag.readCount = 1
	_tag.xtype = tp
	return _tag
}

func (t *Tag) GetValue() interface{} {
	reader := bytes.NewReader(t.value)

	switch t.xtype {
	case NULL:
		return nil
	case SINT:
		result := int8(0)
		lib.ReadByte(reader, &result)
		return result
	case INT:
		result := int16(0)
		lib.ReadByte(reader, &result)
		return result
	case DINT:
		result := int32(0)
		lib.ReadByte(reader, &result)
		return result
	case LINT:
		result := int64(0)
		lib.ReadByte(reader, &result)
		return result
	default:
		return nil
	}
}

func (t *Tag) SetValue(data interface{}) {
	writer := new(bytes.Buffer)

	switch t.xtype {
	case NULL:
	case SINT:
		result := data.(int8)
		lib.WriteByte(writer, &result)
	case INT:
		result := data.(int16)
		lib.WriteByte(writer, &result)
	case DINT:
		result := data.(int32)
		lib.WriteByte(writer, &result)
	case LINT:
		result := data.(int64)
		lib.WriteByte(writer, &result)
	}

	if writer.Len() > 0 {
		t.value = writer.Bytes()
	}
}
