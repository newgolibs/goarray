package goarray

type ArrInterface interface {
    /**    获取指定的索引值    */
    GetbyIndex(index int)interface{}
    /**    切片转换成json字符串    */
    Tojson()string
    /* 属性[first] get需要重写 */
    Getfirst() interface{}    /* 属性[last] get需要重写 */
    Getlast() interface{}    /* 属性[Value] set需要重写 */
    SetValue(Value interface{}) *Arr
}

/**
提供切片的一些基础功能;
*/
type Arr struct
{
    /*取出最前面的元素*/
    first interface{}
    /*取出最后面的元素*/
    last interface{}
    /*切片的内容*/
    Value interface{}
}



    //set -
    func (this *Arr) Setfirst(first interface{}) *Arr{
        this.first = first;
        return this
    }

    //set -
    func (this *Arr) Setlast(last interface{}) *Arr{
        this.last = last;
        return this
    }

    //get -
    func (this *Arr) GetValue() interface{}{
        return this.Value;
    }




//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ ArrInterface =new(Arr)
