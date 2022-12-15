import time,sys
from collections import deque
start=time.time()

inf = sys.argv[1] if len(sys.argv) > 1 else 'inputs/d13distresssignal.txt'
with open(inf) as fi:
    pairs =  [x.split('\n') for x in fi.read()[:-1].split('\n\n')]
def findclosingbracket(a):
    z = 1
    for pos,x in enumerate(a[1:]):
        if x =='[': z+=1
        elif x ==']': z-=1
        if z==0: return pos+1

def compare(a,b):
    print(a,b)
    result = -1
    while a != '' and b != '' and result == -1:
        #print('a=',a,'b=',b,'result',result)
        if a[0] == '[' or b[0] == '[':
            closingbracketa, closingbracketb = -1,-1
            if a[0] == '[' and b[0] != '[':
                closingbracketa =findclosingbracket(a)
                suba = a[1:closingbracketa]
                subb = b[:b.find(',')] if b.find(',')>0 else b
            elif a[0] != '[' and b[0] == '[':
                closingbracketb =findclosingbracket(b)
                suba = a[:a.find(',')] if a.find(',')>0 else a
                subb = b[1:closingbracketb]
            elif a[0] == '[' and b[0] == '[':
                closingbracketa =findclosingbracket(a)
                closingbracketb =findclosingbracket(b)
                suba = a[1:closingbracketa]
                subb = b[1:closingbracketb]
            result = compare(suba,subb)
            if result != -1: return result
            a = a[closingbracketa+2:] if closingbracketa >-1 else a[2:]
            if len(a)>0:
                if a[0] == ',': a = a[1:]
            b = b[closingbracketb+2:] if closingbracketb >-1 else b[2:]
            if len(b)>0:
                if b[0] == ',': b = b[1:]
        elif a[0] in '012345678910' and b[0] in '012345678910':
            xa = int(a[:a.find(',')]) if a.find(',')>0 else int(a)
            xb = int(b[:b.find(',')]) if b.find(',')>0 else int(b)
            if xa>xb: return 0
            elif xa<xb: return 1
            a = a[a.find(',')+1:] if a.find(',') > 0  else ''
            b = b[b.find(',')+1:] if b.find(',') > 0 else  ''
    if len(a)==0  and len(b)!=0:
            return 1
    if len(a)>0 and len(b)==0:
            return 0
    return result

i = 0
count = 0
packets = ['[[2]]','[[6]]']
for pair in pairs:
    i += 1
    a,b = pair
    print(i, compare(a,b) == 1)
    count += compare(a,b)*i
print('part1', count)
"""
for pair in pairs:
    for a in pair:
        for pos,x in enumerate(packets):
            if compare(a,x)==1:

                if a not in packets: packets.insert(pos,a)
        if a not in packets:
                packets.append(a)
print('part2', (packets.index('[[2]]')+1)*(packets.index('[[6]]')+1))

end=time.time()
print(round(end-start,6))"""