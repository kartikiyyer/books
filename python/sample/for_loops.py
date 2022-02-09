
def for_loop():
	for i in range(0, 10):
		print 'Hello World! %s' % i

	for i in range(0, len('Kartik')):
		print '%s' % i

	age = [1,2,3,4]
	for i in age:
		print '%s' % (i)
		
if __name__=='__main__':
	for_loop()
