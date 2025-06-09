# backend

- Selvitä miten youtubesta saa tekstitykset api:N kautta.

## Database

- saved_words taulu johon tallennetaan
   - id, sana, käännös, learning level (vaikka 1-4)
- lessons taulu jossa on:
   - id, blob jossa säilötään teksti?

# frontend

- etusivulla näkyy tallennetut lessonit jossa on uusia sanoja.
- all lessions sivu jossa näkyy kaikki. ja muokkaus?
- Mahdollisuus upata joku teksti/tekstitys tiedostoja?

# Display

- lause on p elementti, sen sisällä jokainen sana on oma span, ja näille spanneille voidaan
lisätä classin avulla tyylitys
- ruudun yläreuna on video embed. Sen alla vierekkäin teksti ja klikatun sanan tiedot/sanakirja
