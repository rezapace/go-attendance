-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 25 Bulan Mei 2023 pada 15.21
-- Versi server: 10.4.25-MariaDB
-- Versi PHP: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `caps`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `dosens`
--

CREATE TABLE `dosens` (
  `id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `nama` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `n_ip` varchar(255) NOT NULL,
  `image` varchar(255) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `dosens`
--

INSERT INTO `dosens` (`id`, `created_at`, `updated_at`, `deleted_at`, `nama`, `email`, `n_ip`, `image`, `user_id`) VALUES
(6, '2023-05-24 10:46:07', '2023-05-24 10:48:07', '2023-05-24 10:48:07', 'John Doe', 'johndoe@example.com', '1234567890', 'https://example.com/image.jpg', 1),
(7, '2023-05-24 10:47:47', '2023-05-24 10:49:12', NULL, 'reza', 'johndoe@example.com', '1234567890', 'https://example.com/image.jpg', 1),
(8, '2023-05-24 10:48:35', '2023-05-24 10:48:35', NULL, 'John Doe', 'johndoe@example.com', '1234567890', 'https://example.com/image.jpg', 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `mahasiswas`
--

CREATE TABLE `mahasiswas` (
  `id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `email` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `nim` varchar(255) NOT NULL,
  `image` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `jurusan` varchar(255) DEFAULT NULL,
  `tahun_masuk` varchar(255) DEFAULT NULL,
  `ip_k` varchar(255) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `mahasiswas`
--

INSERT INTO `mahasiswas` (`id`, `created_at`, `updated_at`, `deleted_at`, `email`, `name`, `nim`, `image`, `phone`, `jurusan`, `tahun_masuk`, `ip_k`, `user_id`) VALUES
(1, '2023-05-25 13:20:38', '2023-05-25 13:21:03', NULL, 'johndoe@example.com', 'dafa', '123456789', 'profile.jpg', '1234567890', 'Teknik Informatika', '2021', '3.75', 1),
(2, '2023-05-25 13:20:40', '2023-05-25 13:20:40', NULL, 'johndoe@example.com', 'daaffa', '123456789', 'profile.jpg', '1234567890', 'Teknik Informatika', '2021', '3.75', 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `email`, `password`) VALUES
(1, '2023-05-01 10:00:00.000', '2023-05-01 10:00:00.000', NULL, 'John Doe', 'johndoe@example.com', 'password123'),
(2, '2023-05-02 11:00:00.000', '2023-05-02 11:00:00.000', NULL, 'Jane Smith', 'janesmith@example.com', 'secret');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `dosens`
--
ALTER TABLE `dosens`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indeks untuk tabel `mahasiswas`
--
ALTER TABLE `mahasiswas`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `dosens`
--
ALTER TABLE `dosens`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT untuk tabel `mahasiswas`
--
ALTER TABLE `mahasiswas`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `dosens`
--
ALTER TABLE `dosens`
  ADD CONSTRAINT `dosens_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
