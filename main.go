package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"nothing.com/scratch/proto/spex/gen/go/voucher_mp_loader.pb"
)

func main() {
	var reqMeta *voucher_mp_loader.RequestMeta
	var reqMetaV2 *voucher_mp_loader.RequestMetaV2
	// rst := CopyMessage(reqMetaV2, reqMeta)
	// fmt.Println("a) nil -> nil")
	// fmt.Printf("   rst -> %T > %v\n", rst, rst)

	// reqMetaV2 = &voucher_mp_loader.RequestMetaV2{}
	// fmt.Println("---")
	// fmt.Println("b) nil -> empty")
	// rst = CopyMessage(reqMetaV2, reqMeta)
	// fmt.Printf("   rst -> %T > %v\n", rst, rst)
	// fmt.Printf("   %T -> %v -> %p\n",
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// )

	// reqMeta = &voucher_mp_loader.RequestMeta{}
	// reqMetaV2 = nil
	// fmt.Println("---")
	// fmt.Println("c) empty -> nil")
	// rst = CopyMessage(reqMetaV2, reqMeta)
	// fmt.Printf("   rst -> %T > %v\n", rst, rst)
	// fmt.Printf("   %T -> %v -> %p\n",
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// )

	// reqMeta = &voucher_mp_loader.RequestMeta{}
	// _ = faker.FakeData(reqMeta,
	// 	options.WithRandomMapAndSliceMaxSize(3),
	// 	options.WithRandomStringLength(5),
	// )
	// reqMetaV2 = &voucher_mp_loader.RequestMetaV2{}
	// fmt.Println("---")
	// fmt.Println("d) some -> empty (*common case)")
	// rst = CopyMessage(reqMetaV2, reqMeta)
	// fmt.Printf("   rst -> %T > %v\n", rst, rst)
	// fmt.Printf("   %T -> %v -> %p\n",
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// )

	reqMeta = &voucher_mp_loader.RequestMeta{}
	_ = faker.FakeData(reqMeta,
		options.WithRandomMapAndSliceMaxSize(5),
		options.WithRandomStringLength(5),
	)
	// fmt.Println("buf -> ", string(buf))
	// reqMetaV2 = &voucher_mp_loader.RequestMetaV2{}
	// fmt.Println("---")
	// fmt.Println("e) some -> empty (*with nil list of int)")
	// rst := CopyMessage(reqMetaV2, reqMeta)
	// fmt.Printf("   rst -> %T > %v; ints > %v(nil?%v)\n",
	// 	rst, rst, rst.ResponseMeta.InvalidIds, rst.ResponseMeta.InvalidIds == nil,
	// )
	// fmt.Printf("   %T -> %v -> %p\n",
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// )

	// reqMeta.GetResponseMeta().InvalidIds = []int32{}
	// buf, _ = proto.Marshal(reqMeta)
	// fmt.Println("buf -> ", string(buf))
	// reqMetaV2 = &voucher_mp_loader.RequestMetaV2{}
	// fmt.Println("---")
	// fmt.Println("f) some -> empty (*with [] list of int)")
	// rst = CopyMessage(reqMetaV2, reqMeta)
	// fmt.Printf("   rst -> %T > %v; ints > %v(nil?%v)\n",
	// 	rst, rst, rst.ResponseMeta.InvalidIds, rst.ResponseMeta.InvalidIds == nil,
	// )
	// fmt.Printf("   %T -> %v -> %p\n",
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// )

	// reqMeta.GetResponseMeta().InvalidIds = []int32{1, 2, 3}
	// buf, _ = proto.Marshal(reqMeta)
	// fmt.Println("buf -> ", string(buf))
	// reqMetaV2 = &voucher_mp_loader.RequestMetaV2{}
	// fmt.Println("---")
	// fmt.Println("g) some -> some (*with [1, 2, 3] list of int)")
	// rst = CopyMessage(reqMetaV2, reqMeta)
	// fmt.Printf("   rst -> %T > %v; ints > %v(nil?%v)\n",
	// 	rst, rst, rst.ResponseMeta.InvalidIds, rst.ResponseMeta.InvalidIds == nil,
	// )
	// fmt.Printf("   %T -> %v -> %p\n",
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// )

	// reqMeta.GetResponseMeta().InvalidIds = []int32{1, 2, 3}
	// buf, _ = proto.Marshal(reqMeta)
	// fmt.Println("buf -> ", string(buf))
	// reqMetaV2 = &voucher_mp_loader.RequestMetaV2{}
	// fmt.Println("---")
	// fmt.Println("h) some -> some (*with [1, 2, 3] list of int)")
	// rst = CopyMessage(reqMetaV2, reqMeta)
	// fmt.Printf("   rst -> %T > %v; ints > %v(nil?%v)\n",
	// 	rst, rst, rst.ResponseMeta.InvalidIds, rst.ResponseMeta.InvalidIds == nil,
	// )
	// fmt.Printf("   %T -> %v -> %p\n",
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// 	rst.GetResponseMeta().GetInvalidIds(),
	// )

	// reqMeta.Cookies = []*voucher_mp_loader.Cookie{
	// 	{
	// 		Name:  proto.String("name"),
	// 		Value: nil,
	// 	},
	// 	{
	// 		Name:  proto.String("again"),
	// 		Value: proto.String("boy"),
	// 	},
	// }
	buf, _ := json.MarshalIndent(reqMeta, "", "  ")
	fmt.Println("buf -> ", string(buf))
	fmt.Println("-----")
	fmt.Println("h) some -> some with nested message")
	reqMetaV2 = &voucher_mp_loader.RequestMetaV2{}
	rst := CopyMessageV2(reqMetaV2, reqMeta)
	buf, _ = json.MarshalIndent(rst, "", "  ")
	fmt.Println("rst -> ", string(buf))
	fmt.Println("-----")
}

func CopyMessageV2[T proto.Message](schema T, source proto.Message) T {
	srcRef := proto.MessageReflect(source)
	// very edge case, when pass in src is a nil
	if !srcRef.IsValid() {
		rst := new(T)
		return *rst
	}
	dstRef := proto.MessageReflect(schema).New()

	sFields := srcRef.Descriptor().Fields()
	dFields := dstRef.Descriptor().Fields()
	for i := 0; i < sFields.Len(); i++ {
		sField := sFields.Get(i)
		dField := dFields.ByName(sField.Name())
		if dField == nil || sField.Kind() != dField.Kind() {
			continue
		}

		switch sField.Kind() {
		case protoreflect.MessageKind:
			if sField.IsList() {
				sList := srcRef.Get(sField).List()
				if !sList.IsValid() {
					continue
				}
				dList := dstRef.Mutable(dField).List()
				for j := 0; j < sList.Len(); j++ {
					dListPb := proto.MessageV1(dList.NewElement().Message())
					sListPb := proto.MessageV1(sList.Get(j).Message())
					child := CopyMessageV2(dListPb, sListPb)
					dList.Append(protoreflect.ValueOf(proto.MessageReflect(child)))
				}
				dstRef.Set(dField, protoreflect.ValueOf(dList))
			} else {
				sMsg := srcRef.Get(sField).Message()
				if !sMsg.IsValid() {
					continue
				}
				dMsg := dstRef.Get(dField).Message()
				child := CopyMessageV2(proto.MessageV1(dMsg), proto.MessageV1(sMsg))
				dstRef.Set(dField, protoreflect.ValueOf(proto.MessageReflect(child)))
			}
		case protoreflect.StringKind, protoreflect.BoolKind,
			protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind,
			protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind, protoreflect.FloatKind,
			protoreflect.Sfixed64Kind, protoreflect.Fixed64Kind, protoreflect.DoubleKind:
			if sField.IsList() {
				sList := srcRef.Get(sField).List()
				if !sList.IsValid() {
					continue
				}
				dValue := dstRef.Mutable(dField)
				dList := dValue.List()
				for j := 0; j < sList.Len(); j++ {
					dList.Append(sList.Get(j))
				}
				dstRef.Set(dField, dValue)
			} else {
				dstRef.Set(dField, srcRef.Get(sField))
			}
		case protoreflect.EnumKind, protoreflect.BytesKind, protoreflect.GroupKind:
			// PENDING
		}
	}
	return proto.MessageV1(dstRef).(T)
}

func CopyMessage(dst proto.Message, src proto.Message, excludes []string) proto.Message {
	dstVal := reflect.ValueOf(dst)
	if dstVal.IsNil() {
		return dst
	}

	srcVal := reflect.ValueOf(src)
	if srcVal.IsNil() {
		return reflect.Zero(dstVal.Type()).Interface().(proto.Message)
	}

	sv := reflect.Indirect(srcVal)
	st := sv.Type()
	sprops := proto.GetProperties(st)

	srcMap := map[string]*reflect.Value{}
	for i := 0; i < sv.NumField(); i++ {
		fv := sv.Field(i)
		if fv.IsValid() {
			props := sprops.Prop[i]
			srcMap[props.OrigName] = &fv
		}
	}

	for _, exclude := range excludes {
		delete(srcMap, exclude)
	}

	dv := reflect.Indirect(dstVal)
	dt := dv.Type()
	dprops := proto.GetProperties(dt)
	for i := 0; i < dv.NumField(); i++ {
		dfv := dv.Field(i)
		dft := dfv.Type().Elem()
		props := dprops.Prop[i]
		if sfv, ok := srcMap[props.OrigName]; ok {
			// same type
			if dfv.Type() == sfv.Type() {
				dfv.Set(*sfv)
				continue
			}

			// int32 to int64
			if dft.Kind() == reflect.Int64 && reflect.Indirect(*sfv).Kind() == reflect.Int32 {
				i := reflect.Indirect(*sfv).Int()
				dfv.Set(reflect.ValueOf(&i))
				continue
			}

			// int32 to uint32
			if dft.Kind() == reflect.Uint32 && reflect.Indirect(*sfv).Kind() == reflect.Int32 {
				ui := uint32(reflect.Indirect(*sfv).Int())
				dfv.Set(reflect.ValueOf(&ui))
				continue
			}

			// copy slice of message
			if dfv.Kind() == reflect.Slice && sfv.Kind() == reflect.Slice {
				dstS := dfv.Interface()
				if dfv.IsNil() {
					dstS = reflect.MakeSlice(dfv.Type(), 0, sfv.Len()).Interface()
				}
				dstS = CopyMessages(dstS, sfv.Interface(), excludes)
				dstV := reflect.ValueOf(dstS)
				if !dstV.IsNil() && dstV.Len() > 0 {
					dfv.Set(dstV)
				}
				continue
			}

			// copy nested message
			if dfv.Kind() != reflect.Ptr || sfv.Kind() != reflect.Ptr {
				continue
			}
			srcM, ok := sfv.Interface().(proto.Message)
			if !ok || sfv.IsNil() {
				continue
			}
			dstM, ok := dfv.Interface().(proto.Message)
			if !ok {
				continue
			}
			if dfv.IsNil() {
				dstM = reflect.New(dfv.Type().Elem()).Interface().(proto.Message)
			}
			dstM = CopyMessage(dstM, srcM, excludes)
			dfv.Set(reflect.ValueOf(dstM))
		}
	}

	return dstVal.Interface().(proto.Message)
}

func CopyMessages(dst interface{}, src interface{}, excludes []string) interface{} {
	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Slice || dstVal.IsNil() {
		return dst
	}

	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() != reflect.Slice || srcVal.IsNil() {
		dst = reflect.Indirect(reflect.New(dstVal.Type())).Interface()
		return dst
	}

	if srcVal.Len() == 0 {
		return dst
	}

	var arrSrc []proto.Message
	for i := 0; i < srcVal.Len(); i++ {
		sv := srcVal.Index(i)
		if !sv.CanInterface() {
			continue
		}
		if v, ok := sv.Interface().(proto.Message); ok {
			arrSrc = append(arrSrc, v)
		}
	}

	dstType := dstVal.Type().Elem()
	if dstType.Kind() != reflect.Ptr {
		return dst
	}
	if !dstType.Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
		return dst
	}
	dstType = dstType.Elem()

	for i := 0; i < len(arrSrc); i++ {
		dstM := reflect.New(dstType).Interface().(proto.Message) //nolint:errcheck
		dstM = CopyMessage(dstM, arrSrc[i], excludes)            //nolint:errcheck
		dstVal = reflect.Append(dstVal, reflect.ValueOf(dstM))
	}

	return dstVal.Interface()
}
