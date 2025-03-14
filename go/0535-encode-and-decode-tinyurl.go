type Codec struct {
    encodeMap map[string]string
    decodeMap map[string]string
}
const base string = "http://tinyurl.com/"

func Constructor() Codec {
    return Codec{make(map[string]string), make(map[string]string)}
}

func (this *Codec) encode(longUrl string) string {
    if _, ok := this.encodeMap[longUrl]; !ok {
        shortUrl := base + strconv.Itoa(len(this.encodeMap) + 1)
        this.encodeMap[longUrl] = shortUrl
        this.decodeMap[shortUrl] = longUrl
    }
    return this.encodeMap[longUrl]
}

func (this *Codec) decode(shortUrl string) string {
    return this.decodeMap[shortUrl]
}
