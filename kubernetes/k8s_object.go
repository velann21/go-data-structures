package main

import (
	"fmt"
	"io/ioutil"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/kubernetes/scheme"
	//"k8s.io/apimachinery/pkg/runtime"
	//"k8s.io/api/core/v1"
)

func main() {
	bytes, err := ioutil.ReadFile("/Users/singaravelannandakumar/go/src/awesomeProject3/kubernetes/dep.yaml")
	if err != nil {
		panic(err.Error())
	}

	decode := scheme.Codecs.UniversalDeserializer().Decode
	rendobj, _, err := decode(bytes, nil, nil)
	if err != nil {
		fmt.Printf("%#v", err)
	}
	//obj := runtime.Object(rendobj)
	fmt.Println(rendobj.GetObjectKind().GroupVersionKind().Kind)
	fmt.Printf("%++v\n\n", rendobj)


	//}

	//var spec v1.Deployment
	//err = yaml.Unmarshal(bytes, &spec)
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Println(spec.Name)
	//fmt.Println(spec.Spec)
}