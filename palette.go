package issa

import (
	"image/color"
)

const StdPaletteBase64Content = "qKioqKlQqKn4qKqgqKtIqKvwqVCoqVFQqVH4qVKgqVNIqVPwqfioqflQqfn4qfqg" +
	"qftIqfvwqqCoqqFQqqH4qqKgqqNIqqPwq0ioq0lQq0n4q0qgq0tIq0vwq/Coq/FQ" +
	"q/H4q/Kgq/NIq/PxUKipUKlRUKn5UKqhUKtJUKvxUVCpUVFRUVH5UVKhUVNJUVPx" +
	"UfipUflRUfn5UfqhUftJUfvxUqCpUqFRUqH5UqKhUqNJUqPxU0ipU0lRU0n5U0qh" +
	"U0tJU0vxU/CpU/FRU/H5U/KhU/NJU/Px+Kip+KlR+Kn5+Kqh+KtJ+Kvx+VCp+VFR" +
	"+VH5+VKh+VNJ+VPx+fip+flR+fn5+fqh+ftJ+fvx+qCp+qFR+qH5+qKh+qNJ+qPx" +
	"+0ip+0lR+0n5+0qh+0tJ+0vx+/Cp+/FR+/H5+/Kh+/NJ+/PyoKiqoKlSoKn6oKqi" +
	"oKtKoKvyoVCqoVFSoVH6oVKioVNKoVPyofiqoflSofn6ofqioftKofvyoqCqoqFS" +
	"oqH6oqKioqNKoqPyo0iqo0lSo0n6o0qio0tKo0vyo/Cqo/FSo/H6o/Kio/NKo/Pz" +
	"SKirSKlTSKn7SKqjSKtLSKvzSVCrSVFTSVH7SVKjSVNLSVPzSfirSflTSfn7Sfqj" +
	"SftLSfvzSqCrSqFTSqH7SqKjSqNLSqPzS0irS0lTS0n7S0qjS0tLS0vzS/CrS/FT" +
	"S/H7S/KjS/NLS/Pz8Kir8KlT8Kn78Kqj8KtL8Kvz8VCr8VFT8VH78VKj8VNL8VPz" +
	"8fir8flT8fn78fqj8ftL8fvz8qCr8qFT8qH78qKj8qNL8qPz80ir80lT80n780qj" +
	"80tL80vz8/Cr8/FT8/H78/Kj8/NL8/PwAAAAAAP8A/wAA////AAD/AP///wD///9" +
	"sbGxsbJZsbMBslmxslpZslsBswGxswJZswMCWbGyWbJaWbMCWlmyWlpaWlsCWwGy" +
	"WwJaWwMDAbGzAbJbAbMDAlmzAlpbAlsDAwGzAwJbAwMAAAAAAAAAAAAAAAAAAAAA" +
	"I/w"

func StdPalette() color.Palette {
	plt := color.Palette{}

	n := uint8(6)
	m := uint8(255) / n // 42

	//adding the first 6 * 6 * 6 = 216 colors in 42 increments
	//from (42,42,42) to (252,252,252)
	for r := uint8(0); r < n; r++ {
		for g := uint8(0); g < n; g++ {
			for b := uint8(0); b < n; b++ {
				c := color.RGBA{
					R: (r + 1) * m,
					G: (g + 1) * m,
					B: (b + 1) * m,
					A: 255,
				}
				plt = append(plt, c)
			}
		}
	}

	//adding 8 more colors from (0,0,0) to (255,255,255)
	//representing combinations of 0 and 255
	for r := uint8(0); r <= 1; r++ {
		for g := uint8(0); g <= 1; g++ {
			for b := uint8(0); b <= 1; b++ {
				c := color.RGBA{
					R: uint8(255) * r,
					G: uint8(255) * g,
					B: uint8(255) * b,
					A: 255,
				}
				plt = append(plt, c)
			}
		}
	}

	max, step := 255-m/2, m

	for r := max - 3*step; r < max; r += step {
		for g := max - 3*step; g < max; g += step {
			for b := max - 3*step; b < max; b += step {
				c := color.RGBA{
					R: r,
					G: g,
					B: b,
					A: 255,
				}
				plt = append(plt, c)
			}
		}
	}

	return plt
}
