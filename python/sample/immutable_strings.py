name = 'Kartik'

# To perform in-place replacement.
name_chars = list(name)
print name_chars
# Replacing a character.
name_chars[1] = 'I'
print name_chars
# New string.
print ''.join(name_chars)
print name

print r'C:\text\new'
s = '''
Hello ''
How  ""
are
you?
'''

print s

print u'sp\xc4m'

print dir(str)
print dir(s)

import pydoc

print pydoc.help(s.decode)