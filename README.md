# Sistem Tarif Parkir

Tugas 2 (Golang) Lumoshive Bootcamp Backend Golang

## Deskripsi

Program untuk menghitung biaya parkir berdasarkan durasi parkir, status member, dan hari libur.

## Aturan Tarif

- Tarif dasar (2 jam pertama): Rp 5.000
- Tarif per jam tambahan (mulai jam ke-3): Rp 2.000
- Biaya tambahan hari libur: Rp 3.000

### Diskon Member

- Member dengan parkir ≤5 jam: diskon 50%
- Member dengan parkir >5 jam: diskon 30%

## Cara Menjalankan

```bash
go run .
```

## Contoh Test Case

```go
// 4 jam, bukan member, bukan hari libur
hitungBiayaParkir(4, false, false) // Output: 9000

// 2 jam, member, hari libur
hitungBiayaParkir(2, true, true) // Output: 4000
```

## Struktur Fungsi

- `cekJamLebih(jamParkir int)` - Menghitung jam parkir lebih dari 2 jam
- `hitungBiayaParkir(jamParkir int, statusMember bool, hariLibur bool)` - Menghitung total biaya parkir
