# XDP (xd protocol)
- Binary

## `Packet`
| 1 byte (version) | 2 bytes (headers length {1} = H) | H bytes (`[]Header`) | 4 bytes (length {2} = N) | N bytes (`[]Payload`) |

## `Payload`
| 2 bytes (UTF-8 key length {3} = K) | K bytes (key) | 1 byte (`DataType Enum`) | 2 bytes (value {4} = L) | L bytes |

## `DataType Enum`

1-3 bits = masks or extended types
4-8 bits = reserved for common types

### 1-3 bits
- 001$_$$$$ - array of **T**

- 0100_0010 - string UTF-8  (array of u16)
- 0110_0011 - string UTF-16 (array of u32)
- 1000_0000 - custom type (would be `any`, `interface{}`)

### 4-8 bits
- 0000_0001 = 1  = u8
- 0000_0010 = 2  = u16
- 0000_0011 = 3  = u32
- 0000_0100 = 4  = u64

- 0000_0101 = 5  = i8
- 0000_0110 = 6  = i16
- 0000_0111 = 7  = i32
- 0000_1000 = 8  = i64

- 0000_1001 = 9  = f32
- 0000_1010 = 10 = f64

- 0000_1011 = 11 = bool

### Notes
- {1} - 64 KB
- {2} - 4 GB
- {3} - max_length = 32768 symbols if char.length == 2 bytes
- {4} - 64 KB
