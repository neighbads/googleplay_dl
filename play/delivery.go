package play

import (
   "41.neocities.org/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

func (a Apk) Url() string {
   value, _ := a.Message.GetBytes(5)()
   return string(value)
}

func (a Apk) Field1() string {
   value, _ := a.Message.GetBytes(1)()
   return string(value)
}

type Apk struct {
   Message protobuf.Message
}

func (d Delivery) Url() string {
   value, _ := d.Message.GetBytes(3)()
   return string(value)
}

type Delivery struct {
   Message protobuf.Message
}

func (d Delivery) Obb() func() (Obb, bool) {
   values := d.Message.Get(4)
   return func() (Obb, bool) {
      value, ok := values()
      return Obb{value}, ok
   }
}

func (d Delivery) Apk() func() (Apk, bool) {
   values := d.Message.Get(15)
   return func() (Apk, bool) {
      value, ok := values()
      return Apk{value}, ok
   }
}

func (g GoogleAuth) Delivery(
   check *GoogleCheckin, app *StoreApp, single bool,
) (*Delivery, error) {
   req, err := http.NewRequest("", "https://android.clients.google.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/fdfe/delivery"
   req.URL.RawQuery = url.Values{
      "doc": {app.Id},
      "vc":  {strconv.FormatUint(app.Version, 10)},
   }.Encode()
   authorization(req, g)
   user_agent(req, single)
   x_dfe_device_id(req, check)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   message := protobuf.Message{}
   err = message.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   message, _ = message.Get(1)()
   message, _ = message.Get(21)()
   switch v, _ := message.GetVarint(1)(); v {
   case 2:
      return nil, errors.New("version")
   case 3:
      return nil, errors.New("acquire")
   }
   message, _ = message.Get(2)()
   return &Delivery{message}, nil
}

func (o Obb) Url() string {
   value, _ := o.Message.GetBytes(4)()
   return string(value)
}

func (o Obb) Field1() uint64 {
   value, _ := o.Message.GetVarint(1)()
   return uint64(value)
}

type Obb struct {
   Message protobuf.Message
}
