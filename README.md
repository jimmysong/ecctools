## Usage

To add two points:

    $ ecctools add <x1> <y1> <x2> <y2>

    $ ecctools add 79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798 483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8 c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5 1ae168fea63dc339a3c58419466ceaeef7f632653266d0e1236431a950cfe52a
    f9308a019258c31049344f85f89d5229b531c845836f99b08601f113bce036f9 388f7b0f632de8140fe337e62a37f3566500a99934c2231b6cb9fd7584b8e672

Results are the new x and y coordinates

To multiply a scalar with a point on the curve:

    $ ecctools mult <k> <x> <y>

    $ ecctools mult 10 79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798 483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8
    e60fce93b59e9ec53011aabc21c23e97b2a31369b87a5ae9c44ee89e2a6dec0a f7e3507399e595929db99f34f57937101296891e44d23f0be1f32cce69616821

To multiply a scalar to the generator point G:

    $ ecctools multg <k>

    $ ecctools multg 10
    e60fce93b59e9ec53011aabc21c23e97b2a31369b87a5ae9c44ee89e2a6dec0a f7e3507399e595929db99f34f57937101296891e44d23f0be1f32cce69616821


Note all numbers on the command line should be in hex.