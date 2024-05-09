# Symulacja Pożaru Lasu

## Przegląd Projektu
Projekt polega na opracowaniu programu, który symuluje pożar lasu wywołany uderzeniem pioruna. Główny koncept zakłada implementację obszaru (np. tablica dwuwymiarowa), gdzie drzewa mogą rosnąć w losowych miejscach. Każde drzewo w lesie jest reprezentowane przez określoną wartość (np. liczba 1) w strukturze danych. Uderzenie pioruna symulowane jest przez wygenerowanie losowych współrzędnych w celu określenia, czy drzewo zostało trafione. Jeśli drzewo zostaje trafione, zaczyna się palić i może zapalić sąsiednie drzewa. Symulacja ma na celu spalenie całego połączonego obszaru drzew, a następnie obliczenie procentowej wartości spalonego lasu.

## Kluczowe Funkcjonalności
1. **Model Pożaru Lasu:** Symulacja uderzeń pioruna i wynikającego z tego rozprzestrzeniania się ognia w oparciu o sąsiedztwo (bok lub róg).
2. **Obliczanie Optymalnej Gęstości Lasu:** Użycie wielokrotnych symulacji losowych do określenia optymalnej gęstości lasu, która balansuje pokrycie drzew z minimalnymi stratami spowodowanymi przez pożary.
3. **Wizualizacja:** Wyświetlanie lasu przed i po pożarze, umożliwiające zarówno proste reprezentacje w terminalu, jak i bardziej zaawansowane graficzne przedstawienia.
4. **Rozszerzalność:** Wprowadzenie dodatkowych parametrów takich jak kierunek wiatru, różnorodna odporność drzew czy kształt lasu, aby zwiększyć złożoność i realizm symulacji.

## Punktacja
- **14 punktów** za symulację pożaru lasu i wyświetlanie procenta spalonych drzew.
- **2 punkty** za obliczenie optymalnego procentu zalesienia za pomocą prób losowych.
- **2 punkty** za wizualizację spalonego lasu.
- **5 punktów** za integrację dodatkowego parametru do modelu pożaru.

## Wymagania Dotyczące Raportu
Raport powinien opisywać część dotyczącą symulacji projektu, szczegółowo przedstawiając przyjęty model pożaru lasu, wyniki symulacji oraz wnioski. Nie jest konieczne opisywanie struktury kodu, jeżeli kod zawiera zwięzłe komentarze wyjaśniające.

## Dodatkowe Uwagi
- Zadanie można podejść na różne sposoby, każde z nich może prowadzić do różnych wyników w zależności od przyjętych założeń i modelu danych.
- Dodanie złożoności takich jak efekty wiatru czy różne cechy drzew znacznie zwiększa trudność zadania. Zaleca się rozpoczęcie od podstawowej wersji i stopniowe wprowadzanie więcej zmiennych.
- Przy rozbudowie zadania należy zachować ostrożność, aby upewnić się, że obliczenia są wykonalne w ramach czasu trwania semestru.

<!-- ## Porady dla Deweloperów
- Zacznij od implementacji podstawowych mechanizmów wzrostu drzew i rozprzestrzeniania się ognia przy użyciu prostej tablicy 2D.
- Upewnij się, że symulacja poprawnie identyfikuje połączone drzewa i zarządza rozprzestrzenianiem się ognia od punktu początkowego.
- Stopniowo wprowadzaj dodatkowe funkcje, takie jak wiatr i wilgotność, po ustaleniu podstawowej funkcjonalności.
- Użyj bibliotek graficznych, jeśli bardziej szczegółowa wizualizacja jest pożądana. -->

