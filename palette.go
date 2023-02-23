package issa

import (
	"image/color"
)

const StdPaletteBase64Content = "qKioqKlQqKn4qKqgqKtIqKvwqVCoqVFQqVH4qVKgqVNIqVPwqfioqflQqfn4qfqg" +
	"qftIqfvwqqCoqqFQqqH4qqKgqqNIqqPwq0ioq0lQq0n4q0qgq0tIq0vwq_Coq_FQq_H4q_Kgq_NIq_PxUKipUKlRUKn5UKqhUKtJUKvxUVCpUVFRUVH5UVKhUVNJUVPxUfipUflRUfn5UfqhUftJUfvxUqCpUqFRUqH5UqKhUqNJUqPxU0ipU0lRU0n5U0qhU0tJU0vxU_CpU_FRU_H5U_KhU_NJU_Px-Kip-KlR-Kn5-Kqh-KtJ-Kvx-VCp-VFR-VH5-VKh-VNJ-VPx-fip-flR-fn5-fqh-ftJ-fvx-qCp-qFR-qH5-qKh-qNJ-qPx-0ip-0lR-0n5-0qh-0tJ-0vx-_Cp-_FR-_H5-_Kh-_NJ-_PyoKiqoKlSoKn6oKqioKtKoKvyoVCqoVFSoVH6oVKioVNKoVPyofiqoflSofn6ofqioftKofvyoqCqoqFSoqH6oqKioqNKoqPyo0iqo0lSo0n6o0qio0tKo0vyo_Cqo_FSo_H6o_Kio_NKo_PzSKirSKlTSKn7SKqjSKtLSKvzSVCrSVFTSVH7SVKjSVNLSVPzSfirSflTSfn7SfqjSftLSfvzSqCrSqFTSqH7SqKjSqNLSqPzS0irS0lTS0n7S0qjS0tLS0vzS_CrS_FTS_H7S_KjS_NLS_Pz8Kir8KlT8Kn78Kqj8KtL8Kvz8VCr8VFT8VH78VKj8VNL8VPz8fir8flT8fn78fqj8ftL8fvz8qCr8qFT8qH78qKj8qNL8qPz80ir80lT80n780qj80tL80vz8_Cr8_FT8_H78_Kj8_NL8_PwAAAAAAP8A_wAA____AAD_AP___wD___9sbGxsbJZsbMBslmxslpZslsBswGxswJZswMCWbGyWbJaWbMCWlmyWlpaWlsCWwGyWwJaWwMDAbGzAbJbAbMDAlmzAlpbAlsDAwGzAwJbAwMAAAAAAAAAAAAAAAAAAAAAI"

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
