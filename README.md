# epoch
Print unix epoch timestamps in human readable form

# Installation

Install via `go install github.com/joedursun/epoch@latest`

# Usage

To convert a single timestamp:

```bash
$ epoch 1685971800000
2023-06-05 09:30:00

$ epoch 1685971800000000
2023-06-05 09:30:00
```

You can also stream timestamps from STDIN using the `-s` flag.

```bash
# timestamps.txt contents:
# 1685221234
# 1686021234567
# 1686124913987654
# 1686224913987432000

$ cat timestamps.txt | epoch -s
2023-05-27 17:00:34
2023-06-05 23:13:54
2023-06-07 04:01:53
2023-06-08 07:48:33
```
