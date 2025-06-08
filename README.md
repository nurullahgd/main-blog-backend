# Blog Backend API

Bu proje, modern bir blog platformu için Go Fiber framework'ü kullanılarak geliştirilmiş bir REST API'dir.

## 🚀 Özellikler

- 🔐 JWT tabanlı kimlik doğrulama
- 👥 Kullanıcı yönetimi
- 📝 Blog yazıları CRUD işlemleri
- 🖼️ Cloudinary entegrasyonu ile resim yükleme
- 🐳 Docker desteği
- 🔄 CORS yapılandırması
- 📊 GORM ile veritabanı işlemleri
- 🛡️ Middleware desteği

## 🛠️ Teknolojiler

- Go 1.x
- Fiber Framework
- GORM
- PostgreSQL
- Docker
- Cloudinary
- JWT

## 📋 Gereksinimler

- Go 1.x veya üzeri
- Docker ve Docker Compose
- PostgreSQL
- Cloudinary hesabı

## 🚀 Kurulum

1. Projeyi klonlayın:
```bash
git clone https://github.com/nurullahgd/main-blog-backend.git
cd main-blog-backend
```

2. Gerekli bağımlılıkları yükleyin:
```bash
go mod download
```

3. `.env` dosyasını oluşturun ve gerekli değişkenleri ayarlayın:
```env
DB_HOST=localhost
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_PORT=5432
JWT_SECRET=your_jwt_secret
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret
```

4. Docker ile çalıştırın:
```bash
docker-compose up -d
```

## 📁 Proje Yapısı

```
.
├── controllers/     # HTTP isteklerini yöneten controller'lar
├── models/         # Veritabanı modelleri
├── routes/         # API rotaları
├── middleware/     # Middleware fonksiyonları
├── utils/          # Yardımcı fonksiyonlar
├── helpers/        # Yardımcı paketler
├── database/       # Veritabanı bağlantı ve konfigürasyon
├── uploads/        # Yüklenen dosyalar için geçici dizin
├── main.go         # Ana uygulama dosyası
├── Dockerfile      # Docker yapılandırması
└── docker-compose.yml
```

## 🔒 API Endpoints

### Kullanıcı İşlemleri
- `POST /api/auth/register` - Yeni kullanıcı kaydı
- `POST /api/auth/login` - Kullanıcı girişi
- `GET /api/users/profile` - Kullanıcı profili
- `PUT /api/users/profile` - Profil güncelleme

### Blog İşlemleri
- `GET /api/blogs` - Tüm blog yazılarını listele
- `GET /api/blogs/:id` - Belirli bir blog yazısını getir
- `POST /api/blogs` - Yeni blog yazısı oluştur
- `PUT /api/blogs/:id` - Blog yazısını güncelle
- `DELETE /api/blogs/:id` - Blog yazısını sil

## 🧪 Test

Testleri çalıştırmak için:
```bash
go test ./tests/...
```

## 📝 Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Detaylar için [LICENSE](LICENSE) dosyasına bakın.

## 👥 Katkıda Bulunma

1. Bu depoyu fork edin
2. Yeni bir branch oluşturun (`git checkout -b feature/amazing-feature`)
3. Değişikliklerinizi commit edin (`git commit -m 'Add some amazing feature'`)
4. Branch'inizi push edin (`git push origin feature/amazing-feature`)
5. Bir Pull Request oluşturun

## 📞 İletişim

Nurullah Gündoğdu - [@nurullahgd](https://github.com/nurullahgd) (https://github.com/nurullahgd)