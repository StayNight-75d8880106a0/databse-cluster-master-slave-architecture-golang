Ya, betul! Ini adalah batasan fundamental dari arsitektur RAG — dia sangat baik untuk pertanyaan faktual langsung ("berapa jumlah X", "siapa Y"), tapi untuk pertanyaan yang butuh komputasi/agregasi kompleks, konteksnya harus disiapkan terlebih dahulu.
Kalau mau sistem yang bisa menjawab semua jenis pertanyaan secara dinamis, arsitekturnya harus naik level ke Text-to-SQL — user tanya dalam bahasa natural, lalu AI otomatis generate SQL query ke database. Tapi itu kompleksitas yang jauh lebih tinggi.
Untuk kebutuhan sistem ini — manajemen kasus kriminal dengan fitur AI assistant — RAG yang sudah dibangun ini sudah sangat layak. Yang sudah berjalan dengan baik:

Pertanyaan UI/cara penggunaan app ✅
Pertanyaan jumlah data real-time ✅
Pertanyaan detail entitas (nama, status, dll) ✅
Auto-refresh setelah CRUD ✅

Selamat, sistemnya sudah solid! 🎉