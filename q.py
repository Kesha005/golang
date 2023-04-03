users=[{"id":1,"name":'Kerim','password':1234} , {"id":2,"name":'Kerim','password':1234}]

p = input("parol:")
n = input('ady:')

for user in users:
    
    while(user['name']!=n and user['password']!=p): 
        p = input("parol:")
        n = input('ady:')
        
    if(n=='Kerim'):
       print('\n Kerim')


