# example-tool-clockwork-base32

> tools of Clockwork Base32, cli and web

## Development

```shellscript
make clean
make build
make dev
```

## Release

```shellscript
git describe --tags --abbrev=0
git tag v0.0.x
git push origin --tags
```

## References

- [https://github.com/szktty/go-clockwork-base32](https://github.com/szktty/go-clockwork-base32)
- [https://github.com/urfave/cli](https://github.com/urfave/cli)
- [https://github.com/magefile/magel](https://github.com/magefile/mage)
- [Go の WASM がライブラリではなくアプリケーションであること](https://www.kabuku.co.jp/developers/annoying-go-wasm)
- [TinyGo を使って export した関数の引数と返り値に string を使う方法](https://kajindowsxp.com/go-tinygo-webassembly-string/)
- [An essay on the bi-directional exchange of strings between the Wasm module (with TinyGo) and Node.js (with WASI support)](https://www.wasm.builders/k33g_org/an-essay-on-the-bi-directional-exchange-of-strings-between-the-wasm-module-with-tinygo-and-nodejs-with-wasi-support-3i9h)
