const colorPaletteBase64Content =
    "qKioqKlQqKn4qKqgqKtIqKvwqVCoqVFQqVH4qVKgqVNIqVPwqfioqflQqfn4qfqg" +
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
    "WwJaWwMDAbGzAbJbAbMDAlmzAlpbAlsDAwGzAwJbAwMAAAAAAAAAAAAAAAAAAAAA"
const mimeTypeBase64Prefix = "data:image/gif;base64,"
const gif89aBase64Header   = "R0lGODlh" // GIF89a

function logicalScreenImageDescriptors(width, height) {
    return new Uint8Array([
        // Logical Screen Descriptor
        width % 256,  // Logical Screen Width
        width / 256,  // Logical Screen Width
        height % 256, // Logical Screen Height
        height / 256, // Logical Screen Height
        0,                   // <Packed Fields>
        0,                   // Background Color Index
        0,                   // Pixel Aspect Ratio
        // Image Descriptor
        0x2C,                // Image Separator
        0,                   // Image Left Position
        0,                   // Image Left Position
        0,                   // Image Top Position
        0,                   // Image Top Position
        width % 256,  // Image Width
        width / 256,  // Image Width
        height % 256, // Image Height
        height / 256, // Image Height
        0x87,                // <Packed Fields>
    ])
}

function colorPalettePrefix(w, h) {
    //https://stackoverflow.com/questions/12710001/how-to-convert-uint8-array-to-base64-encoded-string#comment73376483_12713326
    let str = String.fromCharCode.apply(null, logicalScreenImageDescriptors(w, h))
    let b64s = btoa(str)
    //https://stackoverflow.com/a/7139207
    b64s = b64s.replace(/\+/g, '-').replace(/\//g, '_').replace(/\=+$/, '')
    return gif89aBase64Header + b64s + colorPaletteBase64Content
}

function hydrateColor(dehydrated) {
    if (dehydrated === "") {
        return ""
    }

    let parts = dehydrated.split("=")
    if (parts.length !== 3) {
        return ""
    }

    let width = parseInt(parts[0])
    let height = parseInt(parts[1])

    let hydratedImage = ""

    if (parts[2].startsWith(gif89aBase64Header)) {
        hydratedImage = mimeTypeBase64Prefix + parts[2]
    } else {
        hydratedImage = mimeTypeBase64Prefix + colorPalettePrefix(width, height) + parts[2]
    }

    return hydratedImage
}
