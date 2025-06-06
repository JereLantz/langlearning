# backend

- Selvitä miten youtubesta saa tekstitykset api:N kautta.

# frontend

- etusivulla näkyy tallennetut lessonit jossa on uusia sanoja.
- all lessions sivu jossa näkyy kaikki. ja muokkaus?

## Database

- saved_words taulu johon tallennetaan
   - id, sana, käännös, learning level (vaikka 1-4)
- lessons taulu jossa on:
   - id, blob jossa säilötään teksti?

# Display

- lause on p elementti, sen sisällä jokainen sana on oma span, ja näille spanneille voidaan
lisätä classin avulla tyylitys
- ruudun yläreuna on video embed. Sen alla vierekkäin teksti ja klikatun sanan tiedot/sanakirja
