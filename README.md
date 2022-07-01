# issa

issa is an image summary string module that produces a small and very compact representation of a source image. Image summaries produced by issa are expected to be in 0.2 - 0.6% of the source image size at default sampling level.

This makes issa image summaries suitable for using as a rich image previews that can be embedded with JSON/HTML payload before actual images are requested. Initial analysis suggested that for a 400Kb image, issa would produce a 3Kb base64 encoded summary, that would take about 2Kb storage space dehydrated. Hydration doesn't take a lot of computational resources and can be done on the critical path. 

Space savings are achieved by two factors:

- issa uses a standard 255 palette and maps all source image colors to that palette. Combined with sampling this is (obviously) the main source of the reduction  
- Additionally, having a standard palette and using GIF image allows issa to remove a reproducible header of the resulting base64 string, that is about 20% of the typical GIF image

## Using issa

Typical workflow of an app using issa has the following steps:

- Summarize, dehydrate and store source images into an image summary
- Dehydrate images as needed and embed into HTML/JSON payload

### Summarizing, dehydrating and storing image summaries

TBD

### Dehydrating images

TBD