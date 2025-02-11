
# python wrapper for package github.com/ashupednekar/testgopy/pycall within overall package pycall
# This is what you import to use the package.
# File is generated by gopy. Do not edit.
# gopy build -output=out -vm=python3 github.com/ashupednekar/testgopy/pycall

# the following is required to enable dlopen to open the _go.so file
import os,sys,inspect,collections
try:
	import collections.abc as _collections_abc
except ImportError:
	_collections_abc = collections

cwd = os.getcwd()
currentdir = os.path.dirname(os.path.abspath(inspect.getfile(inspect.currentframe())))
os.chdir(currentdir)
from . import _pycall
from . import go

os.chdir(cwd)

# to use this code in your end-user python file, import it as follows:
# from pycall import pycall
# and then refer to everything using pycall. prefix
# packages imported by this package listed below:




# ---- Types ---


#---- Enums from Go (collections of consts with same type) ---


#---- Constants from Go: Python can only ask that you please don't change these! ---


# ---- Global Variables: can only use functions to access ---


# ---- Interfaces ---


# ---- Structs ---

# Python type for struct pycall.Caller
class Caller(go.GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameters can be unnamed in order of field names or named fields
		in which case a new Go object is constructed first
		"""
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
			_pycall.IncRef(self.handle)
		elif len(args) == 1 and isinstance(args[0], go.GoClass):
			self.handle = args[0].handle
			_pycall.IncRef(self.handle)
		else:
			self.handle = _pycall.pycall_Caller_CTor()
			_pycall.IncRef(self.handle)
			if  1 < len(args):
				self.Iter = args[1]
			if "Iter" in kwargs:
				self.Iter = kwargs["Iter"]
	def __del__(self):
		_pycall.DecRef(self.handle)
	def __str__(self):
		pr = [(p, getattr(self, p)) for p in dir(self) if not p.startswith('__')]
		sv = 'pycall.Caller{'
		first = True
		for v in pr:
			if callable(v[1]):
				continue
			if first:
				first = False
			else:
				sv += ', '
			sv += v[0] + '=' + str(v[1])
		return sv + '}'
	def __repr__(self):
		pr = [(p, getattr(self, p)) for p in dir(self) if not p.startswith('__')]
		sv = 'pycall.Caller ( '
		for v in pr:
			if not callable(v[1]):
				sv += v[0] + '=' + str(v[1]) + ', '
		return sv + ')'
	@property
	def Iter(self):
		return _pycall.pycall_Caller_Iter_Get(self.handle)
	@Iter.setter
	def Iter(self, value):
		if isinstance(value, go.GoClass):
			_pycall.pycall_Caller_Iter_Set(self.handle, value.handle)
		else:
			_pycall.pycall_Caller_Iter_Set(self.handle, value)
	def Start(self, f, goRun=False):
		"""Start(callable f) """
		_pycall.pycall_Caller_Start(self.handle, f, goRun)


# ---- Slices ---


# ---- Maps ---


# ---- Constructors ---
def NewCaller(n):
	"""NewCaller(int n) object"""
	return Caller(handle=_pycall.pycall_NewCaller(n))


# ---- Functions ---


