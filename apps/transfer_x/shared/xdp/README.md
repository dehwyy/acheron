# XDP (xd protocol)
- Binary

## `Packet`
| Name                    | Length (bytes) | Description            |
| ----------------------- | ---------------| ---------------------- |
| Version                 | 1              |                        |
| PacketType              | 1              | `PacketType`           |
| HeadersLength           | 2 -> H         | Length of `Headers`    |
| PayloadLength           | 4 -> N         | Length of `Payload`    |
| Headers                 | H              | `Headers` = `[]Header` |
| Payload                 | N              | Payload = `[]Field`    |


### `PacketType`
| Name           | Value |
| -------------- | ----- |
| Request        | 0x01  |
| Response       | 0x02  |
| StreamRequest  | 0x03  |
| StreamResponse | 0x04  |

## `Header`
| Name           | Length (bytes) | Description |
| -------------- | -------------- | ----------- |
| `Key` length   | 1 -> K         |             |
| `Value` length | 2 -> V         |             |
| `Key`          | K              | UTF-8       |
| `Value`        | V              | UTF-8       |


## `Field`
| Name           | Length (bytes) | Description |
| -------------- | -------------- | ----------- |
| `Key` length   | 1 -> K         |             |
| `Value` length | 4 -> V         |             |
| `DataType`     | 1              | Data Type   |
| `Key`          | K              | UTF-8       |
| `Value`        | V              |             |

## `DataType Enum`

1-3 bits = masks or extended types
4-8 bits = reserved for common types

### 1-3 bits
- 100$_$$$$ - array of **T**

- 0010_0000 - string UTF-8  (array of u16)
- 0010_0001 - string UTF-16 (array of u32)

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
