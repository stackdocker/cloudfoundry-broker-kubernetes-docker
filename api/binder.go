
package api

import (
    "io"
    "encoding/json"
    "bytes"
    "log"
)

type Binder struct {
    Binding  ServiceBinding
    Bound    *ServiceBound
}

func NewBinder(instanceId, bindingId string) *Binder {
    return & Binder{
        Binding: ServiceBinding{
            Id: bindingId,
            InstanceId: instanceId,
        },
        //Bound: &ServiceBound{}
    }
}

func (b *Binder) config(body io.ReadCloser) error {
    dec := json.NewDecoder(body)
    if err := dec.Decode(&b.Binding); err == io.EOF {
		log.Println(b.Binding)
	} else if err != nil {
		//log.Fatal(err)
		log.Println(err)
		return err
	}
	return nil
}

func (b *Binder) Do(body io.ReadCloser) {
    b.config(body)
    b.Bound = &ServiceBound{}
}

func (b *Binder) Result() (*bytes.Buffer, error) {
    bound, err := json.Marshal(b.Bound)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bytes.NewBuffer(bound), nil
}