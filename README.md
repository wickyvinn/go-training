## notes

### slices
[a:z], where a is inclusive and z is exclusive.
length has to be <= capacity.
the length of a slice is the number of elements it contains.
the capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
b = a[1:] // here it keeps the same array address or something?  <- ask adam about this.

### maps
use map[string]string{} rather than make(map[string]string). it's nicer and it's the same perf.
if you do a lookup and the key doesn't exist, you'll get the zero value. (not a nil/pointer reference, not a panic.)

### defer
called when the function leaves scope

### functions are first-class (what does that mean though)
y := func() int { return 456 }() <- in this example, y is actually of type int

### types
"if you're changing something on an object, you have to write the function on a pointer object."

### pointers
if you reference a pointer, then go puts it on the heap
which makes it available outside of the scope of the function. kind of.
you can do - g := thing{}
and if you need it outside of that scope, you can use &g. where go will be like "oh i better put that on the heap."

### interfaces
