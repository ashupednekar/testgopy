/*
cgo stubs for package pycall.
File is generated by gopy. Do not edit.
gopy build -output=out -vm=python3 github.com/ashupednekar/testgopy/pycall
*/

package main

/*

#cgo CFLAGS: "-I/opt/homebrew/opt/python@3.12/Frameworks/Python.framework/Versions/3.12/include/python3.12" -Wno-error -Wno-implicit-function-declaration -Wno-int-conversion
#cgo LDFLAGS: "-L/opt/homebrew/opt/python@3.12/Frameworks/Python.framework/Versions/3.12/lib" "-lpython3.12" -ldl  -framework CoreFoundation

// #define Py_LIMITED_API // need full API for PyRun*
#include <Python.h>
typedef uint8_t bool;
// static inline is trick for avoiding need for extra .c file
// the following are used for build value -- switch on reflect.Kind
// or the types equivalent
static inline PyObject* gopy_build_bool(uint8_t val) {
	return Py_BuildValue("b", val);
}
static inline PyObject* gopy_build_int64(int64_t val) {
	return Py_BuildValue("k", val);
}
static inline PyObject* gopy_build_uint64(uint64_t val) {
	return Py_BuildValue("K", val);
}
static inline PyObject* gopy_build_float64(double val) {
	return Py_BuildValue("d", val);
}
static inline PyObject* gopy_build_string(const char* val) {
	return Py_BuildValue("s", val);
}
static inline void gopy_decref(PyObject* obj) { // macro
	Py_XDECREF(obj);
}
static inline void gopy_incref(PyObject* obj) { // macro
	Py_XINCREF(obj);
}
static inline int gopy_method_check(PyObject* obj) { // macro
	return PyMethod_Check(obj);
}
static inline void gopy_err_handle() {
	if(PyErr_Occurred() != NULL) {
		PyErr_Print();
	}
}

*/
import "C"
import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/go-python/gopy/gopyh" // handler

	"github.com/ashupednekar/testgopy/pycall"
)

// main doesn't do anything in lib / pkg mode, but is essential for exe mode
func main() {

}

// initialization functions -- can be called from python after library is loaded
// GoPyInitRunFile runs a separate python file -- call in GoPyInit if it
// steals the main thread e.g., for GUI event loop, as in GoGi startup.

//export GoPyInit
func GoPyInit() {

}

// type for the handle -- int64 for speed (can switch to string)
type GoHandle int64
type CGoHandle C.longlong

// DecRef decrements the reference count for the specified handle
// and deletes it it goes to zero.
//
//export DecRef
func DecRef(handle CGoHandle) {
	gopyh.DecRef(gopyh.CGoHandle(handle))
}

// IncRef increments the reference count for the specified handle.
//
//export IncRef
func IncRef(handle CGoHandle) {
	gopyh.IncRef(gopyh.CGoHandle(handle))
}

// NumHandles returns the number of handles currently in use.
//
//export NumHandles
func NumHandles() int {
	return gopyh.NumHandles()
}

// boolGoToPy converts a Go bool to python-compatible C.char
func boolGoToPy(b bool) C.char {
	if b {
		return 1
	}
	return 0
}

// boolPyToGo converts a python-compatible C.Char to Go bool
func boolPyToGo(b C.char) bool {
	if b != 0 {
		return true
	}
	return false
}

func complex64GoToPy(c complex64) *C.PyObject {
	return C.PyComplex_FromDoubles(C.double(real(c)), C.double(imag(c)))
}

func complex64PyToGo(o *C.PyObject) complex64 {
	v := C.PyComplex_AsCComplex(o)
	return complex(float32(v.real), float32(v.imag))
}

func complex128GoToPy(c complex128) *C.PyObject {
	return C.PyComplex_FromDoubles(C.double(real(c)), C.double(imag(c)))
}

func complex128PyToGo(o *C.PyObject) complex128 {
	v := C.PyComplex_AsCComplex(o)
	return complex(float64(v.real), float64(v.imag))
}

// errorGoToPy converts a Go error to python-compatible C.CString
func errorGoToPy(e error) *C.char {
	if e != nil {
		return C.CString(e.Error())
	}
	return C.CString("")
}

// --- generated code for package: pycall below: ---

// ---- External Types Outside of Targeted Packages ---

// ---- Package: go ---

// ---- Types ---

// Converters for implicit pointer handles for type: []bool
func ptrFromHandle_Slice_bool(h CGoHandle) *[]bool {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]bool")
	if p == nil {
		return nil
	}
	return p.(*[]bool)
}
func deptrFromHandle_Slice_bool(h CGoHandle) []bool {
	p := ptrFromHandle_Slice_bool(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_bool(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]bool", p))
}

// --- wrapping slice: []bool ---
//
//export Slice_bool_CTor
func Slice_bool_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_bool(&[]bool{}))
}

//export Slice_bool_len
func Slice_bool_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_bool(handle))
}

//export Slice_bool_elem
func Slice_bool_elem(handle CGoHandle, _idx int) C.char {
	s := deptrFromHandle_Slice_bool(handle)
	return boolGoToPy(s[_idx])
}

//export Slice_bool_subslice
func Slice_bool_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_bool(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_bool(&ss))
}

//export Slice_bool_set
func Slice_bool_set(handle CGoHandle, _idx int, _vl C.char) {
	s := deptrFromHandle_Slice_bool(handle)
	s[_idx] = boolPyToGo(_vl)
}

//export Slice_bool_append
func Slice_bool_append(handle CGoHandle, _vl C.char) {
	s := ptrFromHandle_Slice_bool(handle)
	*s = append(*s, boolPyToGo(_vl))
}

// Converters for implicit pointer handles for type: []byte
func ptrFromHandle_Slice_byte(h CGoHandle) *[]byte {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]byte")
	if p == nil {
		return nil
	}
	return p.(*[]byte)
}
func deptrFromHandle_Slice_byte(h CGoHandle) []byte {
	p := ptrFromHandle_Slice_byte(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_byte(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]byte", p))
}

// --- wrapping slice: []byte ---
//
//export Slice_byte_CTor
func Slice_byte_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_byte(&[]byte{}))
}

//export Slice_byte_len
func Slice_byte_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_byte(handle))
}

//export Slice_byte_elem
func Slice_byte_elem(handle CGoHandle, _idx int) C.char {
	s := deptrFromHandle_Slice_byte(handle)
	return C.char(s[_idx])
}

//export Slice_byte_subslice
func Slice_byte_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_byte(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_byte(&ss))
}

//export Slice_byte_set
func Slice_byte_set(handle CGoHandle, _idx int, _vl C.char) {
	s := deptrFromHandle_Slice_byte(handle)
	s[_idx] = byte(_vl)
}

//export Slice_byte_append
func Slice_byte_append(handle CGoHandle, _vl C.char) {
	s := ptrFromHandle_Slice_byte(handle)
	*s = append(*s, byte(_vl))
}

//export Slice_byte_from_bytes
func Slice_byte_from_bytes(o *C.PyObject) CGoHandle {
	size := C.PyBytes_Size(o)
	ptr := unsafe.Pointer(C.PyBytes_AsString(o))
	data := make([]byte, size)
	tmp := unsafe.Slice((*byte)(ptr), size)
	copy(data, tmp)
	return handleFromPtr_Slice_byte(&data)
}

//export Slice_byte_to_bytes
func Slice_byte_to_bytes(handle CGoHandle) *C.PyObject {
	s := deptrFromHandle_Slice_byte(handle)
	ptr := unsafe.Pointer(&s[0])
	size := len(s)
	return C.PyBytes_FromStringAndSize((*C.char)(ptr), C.long(size))
}

// Converters for implicit pointer handles for type: []error
func ptrFromHandle_Slice_error(h CGoHandle) *[]error {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]error")
	if p == nil {
		return nil
	}
	return p.(*[]error)
}
func deptrFromHandle_Slice_error(h CGoHandle) []error {
	p := ptrFromHandle_Slice_error(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_error(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]error", p))
}

// --- wrapping slice: []error ---
//
//export Slice_error_CTor
func Slice_error_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_error(&[]error{}))
}

//export Slice_error_len
func Slice_error_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_error(handle))
}

//export Slice_error_elem
func Slice_error_elem(handle CGoHandle, _idx int) *C.char {
	s := deptrFromHandle_Slice_error(handle)
	return errorGoToPy(s[_idx])
}

//export Slice_error_subslice
func Slice_error_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_error(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_error(&ss))
}

//export Slice_error_set
func Slice_error_set(handle CGoHandle, _idx int, _vl *C.char) {
	s := deptrFromHandle_Slice_error(handle)
	s[_idx] = errors.New(C.GoString(_vl))
}

//export Slice_error_append
func Slice_error_append(handle CGoHandle, _vl *C.char) {
	s := ptrFromHandle_Slice_error(handle)
	*s = append(*s, errors.New(C.GoString(_vl)))
}

// Converters for implicit pointer handles for type: []float32
func ptrFromHandle_Slice_float32(h CGoHandle) *[]float32 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]float32")
	if p == nil {
		return nil
	}
	return p.(*[]float32)
}
func deptrFromHandle_Slice_float32(h CGoHandle) []float32 {
	p := ptrFromHandle_Slice_float32(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_float32(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]float32", p))
}

// --- wrapping slice: []float32 ---
//
//export Slice_float32_CTor
func Slice_float32_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_float32(&[]float32{}))
}

//export Slice_float32_len
func Slice_float32_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_float32(handle))
}

//export Slice_float32_elem
func Slice_float32_elem(handle CGoHandle, _idx int) C.float {
	s := deptrFromHandle_Slice_float32(handle)
	return C.float(s[_idx])
}

//export Slice_float32_subslice
func Slice_float32_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_float32(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_float32(&ss))
}

//export Slice_float32_set
func Slice_float32_set(handle CGoHandle, _idx int, _vl C.float) {
	s := deptrFromHandle_Slice_float32(handle)
	s[_idx] = float32(_vl)
}

//export Slice_float32_append
func Slice_float32_append(handle CGoHandle, _vl C.float) {
	s := ptrFromHandle_Slice_float32(handle)
	*s = append(*s, float32(_vl))
}

// Converters for implicit pointer handles for type: []float64
func ptrFromHandle_Slice_float64(h CGoHandle) *[]float64 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]float64")
	if p == nil {
		return nil
	}
	return p.(*[]float64)
}
func deptrFromHandle_Slice_float64(h CGoHandle) []float64 {
	p := ptrFromHandle_Slice_float64(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_float64(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]float64", p))
}

// --- wrapping slice: []float64 ---
//
//export Slice_float64_CTor
func Slice_float64_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_float64(&[]float64{}))
}

//export Slice_float64_len
func Slice_float64_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_float64(handle))
}

//export Slice_float64_elem
func Slice_float64_elem(handle CGoHandle, _idx int) C.double {
	s := deptrFromHandle_Slice_float64(handle)
	return C.double(s[_idx])
}

//export Slice_float64_subslice
func Slice_float64_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_float64(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_float64(&ss))
}

//export Slice_float64_set
func Slice_float64_set(handle CGoHandle, _idx int, _vl C.double) {
	s := deptrFromHandle_Slice_float64(handle)
	s[_idx] = float64(_vl)
}

//export Slice_float64_append
func Slice_float64_append(handle CGoHandle, _vl C.double) {
	s := ptrFromHandle_Slice_float64(handle)
	*s = append(*s, float64(_vl))
}

// Converters for implicit pointer handles for type: []int
func ptrFromHandle_Slice_int(h CGoHandle) *[]int {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]int")
	if p == nil {
		return nil
	}
	return p.(*[]int)
}
func deptrFromHandle_Slice_int(h CGoHandle) []int {
	p := ptrFromHandle_Slice_int(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_int(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]int", p))
}

// --- wrapping slice: []int ---
//
//export Slice_int_CTor
func Slice_int_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_int(&[]int{}))
}

//export Slice_int_len
func Slice_int_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_int(handle))
}

//export Slice_int_elem
func Slice_int_elem(handle CGoHandle, _idx int) C.longlong {
	s := deptrFromHandle_Slice_int(handle)
	return C.longlong(s[_idx])
}

//export Slice_int_subslice
func Slice_int_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_int(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_int(&ss))
}

//export Slice_int_set
func Slice_int_set(handle CGoHandle, _idx int, _vl C.longlong) {
	s := deptrFromHandle_Slice_int(handle)
	s[_idx] = int(_vl)
}

//export Slice_int_append
func Slice_int_append(handle CGoHandle, _vl C.longlong) {
	s := ptrFromHandle_Slice_int(handle)
	*s = append(*s, int(_vl))
}

// Converters for implicit pointer handles for type: []int16
func ptrFromHandle_Slice_int16(h CGoHandle) *[]int16 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]int16")
	if p == nil {
		return nil
	}
	return p.(*[]int16)
}
func deptrFromHandle_Slice_int16(h CGoHandle) []int16 {
	p := ptrFromHandle_Slice_int16(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_int16(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]int16", p))
}

// --- wrapping slice: []int16 ---
//
//export Slice_int16_CTor
func Slice_int16_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_int16(&[]int16{}))
}

//export Slice_int16_len
func Slice_int16_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_int16(handle))
}

//export Slice_int16_elem
func Slice_int16_elem(handle CGoHandle, _idx int) C.short {
	s := deptrFromHandle_Slice_int16(handle)
	return C.short(s[_idx])
}

//export Slice_int16_subslice
func Slice_int16_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_int16(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_int16(&ss))
}

//export Slice_int16_set
func Slice_int16_set(handle CGoHandle, _idx int, _vl C.short) {
	s := deptrFromHandle_Slice_int16(handle)
	s[_idx] = int16(_vl)
}

//export Slice_int16_append
func Slice_int16_append(handle CGoHandle, _vl C.short) {
	s := ptrFromHandle_Slice_int16(handle)
	*s = append(*s, int16(_vl))
}

// Converters for implicit pointer handles for type: []int32
func ptrFromHandle_Slice_int32(h CGoHandle) *[]int32 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]int32")
	if p == nil {
		return nil
	}
	return p.(*[]int32)
}
func deptrFromHandle_Slice_int32(h CGoHandle) []int32 {
	p := ptrFromHandle_Slice_int32(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_int32(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]int32", p))
}

// --- wrapping slice: []int32 ---
//
//export Slice_int32_CTor
func Slice_int32_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_int32(&[]int32{}))
}

//export Slice_int32_len
func Slice_int32_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_int32(handle))
}

//export Slice_int32_elem
func Slice_int32_elem(handle CGoHandle, _idx int) C.long {
	s := deptrFromHandle_Slice_int32(handle)
	return C.long(s[_idx])
}

//export Slice_int32_subslice
func Slice_int32_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_int32(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_int32(&ss))
}

//export Slice_int32_set
func Slice_int32_set(handle CGoHandle, _idx int, _vl C.long) {
	s := deptrFromHandle_Slice_int32(handle)
	s[_idx] = int32(_vl)
}

//export Slice_int32_append
func Slice_int32_append(handle CGoHandle, _vl C.long) {
	s := ptrFromHandle_Slice_int32(handle)
	*s = append(*s, int32(_vl))
}

// Converters for implicit pointer handles for type: []int64
func ptrFromHandle_Slice_int64(h CGoHandle) *[]int64 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]int64")
	if p == nil {
		return nil
	}
	return p.(*[]int64)
}
func deptrFromHandle_Slice_int64(h CGoHandle) []int64 {
	p := ptrFromHandle_Slice_int64(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_int64(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]int64", p))
}

// --- wrapping slice: []int64 ---
//
//export Slice_int64_CTor
func Slice_int64_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_int64(&[]int64{}))
}

//export Slice_int64_len
func Slice_int64_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_int64(handle))
}

//export Slice_int64_elem
func Slice_int64_elem(handle CGoHandle, _idx int) C.longlong {
	s := deptrFromHandle_Slice_int64(handle)
	return C.longlong(s[_idx])
}

//export Slice_int64_subslice
func Slice_int64_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_int64(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_int64(&ss))
}

//export Slice_int64_set
func Slice_int64_set(handle CGoHandle, _idx int, _vl C.longlong) {
	s := deptrFromHandle_Slice_int64(handle)
	s[_idx] = int64(_vl)
}

//export Slice_int64_append
func Slice_int64_append(handle CGoHandle, _vl C.longlong) {
	s := ptrFromHandle_Slice_int64(handle)
	*s = append(*s, int64(_vl))
}

// Converters for implicit pointer handles for type: []int8
func ptrFromHandle_Slice_int8(h CGoHandle) *[]int8 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]int8")
	if p == nil {
		return nil
	}
	return p.(*[]int8)
}
func deptrFromHandle_Slice_int8(h CGoHandle) []int8 {
	p := ptrFromHandle_Slice_int8(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_int8(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]int8", p))
}

// --- wrapping slice: []int8 ---
//
//export Slice_int8_CTor
func Slice_int8_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_int8(&[]int8{}))
}

//export Slice_int8_len
func Slice_int8_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_int8(handle))
}

//export Slice_int8_elem
func Slice_int8_elem(handle CGoHandle, _idx int) C.char {
	s := deptrFromHandle_Slice_int8(handle)
	return C.char(s[_idx])
}

//export Slice_int8_subslice
func Slice_int8_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_int8(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_int8(&ss))
}

//export Slice_int8_set
func Slice_int8_set(handle CGoHandle, _idx int, _vl C.char) {
	s := deptrFromHandle_Slice_int8(handle)
	s[_idx] = int8(_vl)
}

//export Slice_int8_append
func Slice_int8_append(handle CGoHandle, _vl C.char) {
	s := ptrFromHandle_Slice_int8(handle)
	*s = append(*s, int8(_vl))
}

// Converters for implicit pointer handles for type: []rune
func ptrFromHandle_Slice_rune(h CGoHandle) *[]rune {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]rune")
	if p == nil {
		return nil
	}
	return p.(*[]rune)
}
func deptrFromHandle_Slice_rune(h CGoHandle) []rune {
	p := ptrFromHandle_Slice_rune(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_rune(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]rune", p))
}

// --- wrapping slice: []rune ---
//
//export Slice_rune_CTor
func Slice_rune_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_rune(&[]rune{}))
}

//export Slice_rune_len
func Slice_rune_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_rune(handle))
}

//export Slice_rune_elem
func Slice_rune_elem(handle CGoHandle, _idx int) C.long {
	s := deptrFromHandle_Slice_rune(handle)
	return C.long(s[_idx])
}

//export Slice_rune_subslice
func Slice_rune_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_rune(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_rune(&ss))
}

//export Slice_rune_set
func Slice_rune_set(handle CGoHandle, _idx int, _vl C.long) {
	s := deptrFromHandle_Slice_rune(handle)
	s[_idx] = rune(_vl)
}

//export Slice_rune_append
func Slice_rune_append(handle CGoHandle, _vl C.long) {
	s := ptrFromHandle_Slice_rune(handle)
	*s = append(*s, rune(_vl))
}

// Converters for implicit pointer handles for type: []string
func ptrFromHandle_Slice_string(h CGoHandle) *[]string {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]string")
	if p == nil {
		return nil
	}
	return p.(*[]string)
}
func deptrFromHandle_Slice_string(h CGoHandle) []string {
	p := ptrFromHandle_Slice_string(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_string(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]string", p))
}

// --- wrapping slice: []string ---
//
//export Slice_string_CTor
func Slice_string_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_string(&[]string{}))
}

//export Slice_string_len
func Slice_string_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_string(handle))
}

//export Slice_string_elem
func Slice_string_elem(handle CGoHandle, _idx int) *C.char {
	s := deptrFromHandle_Slice_string(handle)
	return C.CString(s[_idx])
}

//export Slice_string_subslice
func Slice_string_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_string(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_string(&ss))
}

//export Slice_string_set
func Slice_string_set(handle CGoHandle, _idx int, _vl *C.char) {
	s := deptrFromHandle_Slice_string(handle)
	s[_idx] = C.GoString(_vl)
}

//export Slice_string_append
func Slice_string_append(handle CGoHandle, _vl *C.char) {
	s := ptrFromHandle_Slice_string(handle)
	*s = append(*s, C.GoString(_vl))
}

// Converters for implicit pointer handles for type: []uint
func ptrFromHandle_Slice_uint(h CGoHandle) *[]uint {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]uint")
	if p == nil {
		return nil
	}
	return p.(*[]uint)
}
func deptrFromHandle_Slice_uint(h CGoHandle) []uint {
	p := ptrFromHandle_Slice_uint(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_uint(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]uint", p))
}

// --- wrapping slice: []uint ---
//
//export Slice_uint_CTor
func Slice_uint_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_uint(&[]uint{}))
}

//export Slice_uint_len
func Slice_uint_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_uint(handle))
}

//export Slice_uint_elem
func Slice_uint_elem(handle CGoHandle, _idx int) C.ulonglong {
	s := deptrFromHandle_Slice_uint(handle)
	return C.ulonglong(s[_idx])
}

//export Slice_uint_subslice
func Slice_uint_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_uint(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_uint(&ss))
}

//export Slice_uint_set
func Slice_uint_set(handle CGoHandle, _idx int, _vl C.ulonglong) {
	s := deptrFromHandle_Slice_uint(handle)
	s[_idx] = uint(_vl)
}

//export Slice_uint_append
func Slice_uint_append(handle CGoHandle, _vl C.ulonglong) {
	s := ptrFromHandle_Slice_uint(handle)
	*s = append(*s, uint(_vl))
}

// Converters for implicit pointer handles for type: []uint16
func ptrFromHandle_Slice_uint16(h CGoHandle) *[]uint16 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]uint16")
	if p == nil {
		return nil
	}
	return p.(*[]uint16)
}
func deptrFromHandle_Slice_uint16(h CGoHandle) []uint16 {
	p := ptrFromHandle_Slice_uint16(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_uint16(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]uint16", p))
}

// --- wrapping slice: []uint16 ---
//
//export Slice_uint16_CTor
func Slice_uint16_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_uint16(&[]uint16{}))
}

//export Slice_uint16_len
func Slice_uint16_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_uint16(handle))
}

//export Slice_uint16_elem
func Slice_uint16_elem(handle CGoHandle, _idx int) C.ushort {
	s := deptrFromHandle_Slice_uint16(handle)
	return C.ushort(s[_idx])
}

//export Slice_uint16_subslice
func Slice_uint16_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_uint16(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_uint16(&ss))
}

//export Slice_uint16_set
func Slice_uint16_set(handle CGoHandle, _idx int, _vl C.ushort) {
	s := deptrFromHandle_Slice_uint16(handle)
	s[_idx] = uint16(_vl)
}

//export Slice_uint16_append
func Slice_uint16_append(handle CGoHandle, _vl C.ushort) {
	s := ptrFromHandle_Slice_uint16(handle)
	*s = append(*s, uint16(_vl))
}

// Converters for implicit pointer handles for type: []uint32
func ptrFromHandle_Slice_uint32(h CGoHandle) *[]uint32 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]uint32")
	if p == nil {
		return nil
	}
	return p.(*[]uint32)
}
func deptrFromHandle_Slice_uint32(h CGoHandle) []uint32 {
	p := ptrFromHandle_Slice_uint32(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_uint32(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]uint32", p))
}

// --- wrapping slice: []uint32 ---
//
//export Slice_uint32_CTor
func Slice_uint32_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_uint32(&[]uint32{}))
}

//export Slice_uint32_len
func Slice_uint32_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_uint32(handle))
}

//export Slice_uint32_elem
func Slice_uint32_elem(handle CGoHandle, _idx int) C.ulong {
	s := deptrFromHandle_Slice_uint32(handle)
	return C.ulong(s[_idx])
}

//export Slice_uint32_subslice
func Slice_uint32_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_uint32(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_uint32(&ss))
}

//export Slice_uint32_set
func Slice_uint32_set(handle CGoHandle, _idx int, _vl C.ulong) {
	s := deptrFromHandle_Slice_uint32(handle)
	s[_idx] = uint32(_vl)
}

//export Slice_uint32_append
func Slice_uint32_append(handle CGoHandle, _vl C.ulong) {
	s := ptrFromHandle_Slice_uint32(handle)
	*s = append(*s, uint32(_vl))
}

// Converters for implicit pointer handles for type: []uint64
func ptrFromHandle_Slice_uint64(h CGoHandle) *[]uint64 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]uint64")
	if p == nil {
		return nil
	}
	return p.(*[]uint64)
}
func deptrFromHandle_Slice_uint64(h CGoHandle) []uint64 {
	p := ptrFromHandle_Slice_uint64(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_uint64(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]uint64", p))
}

// --- wrapping slice: []uint64 ---
//
//export Slice_uint64_CTor
func Slice_uint64_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_uint64(&[]uint64{}))
}

//export Slice_uint64_len
func Slice_uint64_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_uint64(handle))
}

//export Slice_uint64_elem
func Slice_uint64_elem(handle CGoHandle, _idx int) C.ulonglong {
	s := deptrFromHandle_Slice_uint64(handle)
	return C.ulonglong(s[_idx])
}

//export Slice_uint64_subslice
func Slice_uint64_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_uint64(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_uint64(&ss))
}

//export Slice_uint64_set
func Slice_uint64_set(handle CGoHandle, _idx int, _vl C.ulonglong) {
	s := deptrFromHandle_Slice_uint64(handle)
	s[_idx] = uint64(_vl)
}

//export Slice_uint64_append
func Slice_uint64_append(handle CGoHandle, _vl C.ulonglong) {
	s := ptrFromHandle_Slice_uint64(handle)
	*s = append(*s, uint64(_vl))
}

// Converters for implicit pointer handles for type: []uint8
func ptrFromHandle_Slice_uint8(h CGoHandle) *[]uint8 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]uint8")
	if p == nil {
		return nil
	}
	return p.(*[]uint8)
}
func deptrFromHandle_Slice_uint8(h CGoHandle) []uint8 {
	p := ptrFromHandle_Slice_uint8(h)
	if p == nil {
		return nil
	}
	return *p
}
func handleFromPtr_Slice_uint8(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]uint8", p))
}

// --- wrapping slice: []uint8 ---
//
//export Slice_uint8_CTor
func Slice_uint8_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_uint8(&[]uint8{}))
}

//export Slice_uint8_len
func Slice_uint8_len(handle CGoHandle) int {
	return len(deptrFromHandle_Slice_uint8(handle))
}

//export Slice_uint8_elem
func Slice_uint8_elem(handle CGoHandle, _idx int) C.uchar {
	s := deptrFromHandle_Slice_uint8(handle)
	return C.uchar(s[_idx])
}

//export Slice_uint8_subslice
func Slice_uint8_subslice(handle CGoHandle, _st, _ed int) CGoHandle {
	s := deptrFromHandle_Slice_uint8(handle)
	ss := s[_st:_ed]
	return CGoHandle(handleFromPtr_Slice_uint8(&ss))
}

//export Slice_uint8_set
func Slice_uint8_set(handle CGoHandle, _idx int, _vl C.uchar) {
	s := deptrFromHandle_Slice_uint8(handle)
	s[_idx] = uint8(_vl)
}

//export Slice_uint8_append
func Slice_uint8_append(handle CGoHandle, _vl C.uchar) {
	s := ptrFromHandle_Slice_uint8(handle)
	*s = append(*s, uint8(_vl))
}

// ---- Package: pycall ---

// ---- Types ---

// Converters for pointer handles for type: *pycall.Caller
func ptrFromHandle_Ptr_pycall_Caller(h CGoHandle) *pycall.Caller {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "*pycall.Caller")
	if p == nil {
		return nil
	}
	return gopyh.Embed(p, reflect.TypeOf(pycall.Caller{})).(*pycall.Caller)
}
func handleFromPtr_Ptr_pycall_Caller(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("*pycall.Caller", p))
}

// Converters for non-pointer handles for type: pycall.Caller
func ptrFromHandle_pycall_Caller(h CGoHandle) *pycall.Caller {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "pycall.Caller")
	if p == nil {
		return nil
	}
	return gopyh.Embed(p, reflect.TypeOf(pycall.Caller{})).(*pycall.Caller)
}
func handleFromPtr_pycall_Caller(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("pycall.Caller", p))
}

// ---- Global Variables: can only use functions to access ---

// ---- Interfaces ---

// ---- Structs ---

// --- wrapping struct: pycall.Caller ---
//
//export pycall_Caller_CTor
func pycall_Caller_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_pycall_Caller(&pycall.Caller{}))
}

//export pycall_Caller_Iter_Get
func pycall_Caller_Iter_Get(handle CGoHandle) C.longlong {
	op := ptrFromHandle_pycall_Caller(handle)
	return C.longlong(op.Iter)
}

//export pycall_Caller_Iter_Set
func pycall_Caller_Iter_Set(handle CGoHandle, val C.longlong) {
	op := ptrFromHandle_pycall_Caller(handle)
	op.Iter = int(val)
}

//export pycall_Caller_Start
func pycall_Caller_Start(_handle CGoHandle, f *C.PyObject, goRun C.char) {
	_fun_arg := f
	_saved_thread := C.PyEval_SaveThread()
	defer C.PyEval_RestoreThread(_saved_thread)
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "*pycall.Caller")
	if __err != nil {
		return
	}
	if boolPyToGo(goRun) {
		go gopyh.Embed(vifc, reflect.TypeOf(pycall.Caller{})).(*pycall.Caller).Start(func() {
			if C.PyCallable_Check(_fun_arg) == 0 {
				return
			}
			_gstate := C.PyGILState_Ensure()
			C.PyObject_CallObject(_fun_arg, nil)
			C.gopy_err_handle()
			C.PyGILState_Release(_gstate)
		})
	} else {
		gopyh.Embed(vifc, reflect.TypeOf(pycall.Caller{})).(*pycall.Caller).Start(func() {
			if C.PyCallable_Check(_fun_arg) == 0 {
				return
			}
			_gstate := C.PyGILState_Ensure()
			C.PyObject_CallObject(_fun_arg, nil)
			C.gopy_err_handle()
			C.PyGILState_Release(_gstate)
		})
	}
}

// ---- Slices ---

// ---- Maps ---

// ---- Constructors ---

//export pycall_NewCaller
func pycall_NewCaller(n C.longlong) CGoHandle {
	_saved_thread := C.PyEval_SaveThread()
	defer C.PyEval_RestoreThread(_saved_thread)
	cret := pycall.NewCaller(int(n))

	return handleFromPtr_pycall_Caller(&cret)
}

// ---- Functions ---
