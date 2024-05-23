Zadanie polega na napisaniu prostego serwera http posiadającego możliwość publikowanie, pobierania i usuwania postów. Treść postów ma pochodzić z pliku Global Shark Attack - World (forma JSON). W zadaniu należy:

Stworzyć strukturę odpowiadającą danym znajdującym się w pliku (można wybrać 6 dowolnych pól)
Wczytać z pliku dane10 losowych rekordów  
 Stworzyć serwer http z funkcjonalnością publikowania, pobierania i usuwania postów
Opublikować na nim posty odpowiadające wczytanym danym
Zademonstrować funkcjonalność pobierania i usuwania postu

Materiały do zadania:

Przykładowy tutorial pracy z plikami JSON: https://tutorialedge.net/golang/parsing-json-with-golang/

(uwaga, io/ioutil ma status depracated, obecnie wczytywanie plików odbywa się przy pomocy os plik można wczytać na przykład przy pomocy os.ReadFile)

Link do prostego serwera (wraz z linkiem do tutorialu): https://github.com/andyjessop/simple-go-server

Link do danych: https://public.opendatasoft.com/explore/dataset/global-shark-attack/export/?disjunctive.country&disjunctive.area&disjunctive.activity

Punktacja:

1. Wczytanie pliku JSON 5pkt

2. Stworzenie serwera 9 pkt

3. Zademonstrowanie funkcjonalności publikacji postów na podstawie stworzonej struktury 3 pkt

4. Zademonstrowanie funkcjonalności pobierania postów 3 pkt

5. Zademonstrowanie funkcjonalności usuwania postów 3 pkt
