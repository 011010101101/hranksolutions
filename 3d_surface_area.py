#!/bin/python3

import sys

def surfaceArea(A):
    # Calculating the surface area of the top and bottom
    # is easy because it's just 2 for every tile. Per
    # the problem definition there can be no holes so
    # we can just multiply.
    top_and_bottom = len(A)*len(A[0])*2
    
    # Here we create a new array B which is the same
    # as A but with zeros wrapping in both dimensions
    #
    # In other words, we change A to 0 0 0
    #                                0 A 0
    #                                0 0 0
    zeros = [0]*(len(A[0])+2)
    B = [zeros]
    for i in A:
        x = [0]
        x.extend(i)
        x.extend([0])
        B.append(x)
    B.append(zeros)
    
    # Finally we loop over the new array B in both
    # directions counting the absolute value of the
    # cell differences.
    
    # In other words 0 2 3 1 0 is evaluated as
    #                 2+1+2+1 = 6.
    
    count = 0
    old = 0
    for x in range(len(A)+2):
        for y in range(len(A[0])+2):
            count += abs(B[x][y] - old)
            old = B[x][y]
        
    for x in range(len(A[0])+2):
        for y in range(len(A)+2):
            count += abs(B[y][x] - old)
            old = B[y][x]
            
    return top_and_bottom+count
