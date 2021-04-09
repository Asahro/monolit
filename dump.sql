-- phpMyAdmin SQL Dump
-- version 4.9.5
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Apr 09, 2021 at 02:27 AM
-- Server version: 10.4.14-MariaDB-cll-lve
-- PHP Version: 7.2.34

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `u877696391_test`
--

-- --------------------------------------------------------

--
-- Table structure for table `data_pegawai`
--

CREATE TABLE `data_pegawai` (
  `id` int(11) NOT NULL,
  `nama_lengkap` varchar(1100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nip` varchar(110) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tempat_lahir` varchar(110) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tanggal_lahir` date NOT NULL,
  `tanggal_bergabung` date NOT NULL,
  `status` int(110) NOT NULL,
  `atasan` int(110) NOT NULL,
  `catatan` text COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `data_pegawai`
--

INSERT INTO `data_pegawai` (`id`, `nama_lengkap`, `nip`, `tempat_lahir`, `tanggal_lahir`, `tanggal_bergabung`, `status`, `atasan`, `catatan`) VALUES
(1, 'Dr Andre Syaefudin', 'RJP880001', 'Serang', '1994-02-26', '2021-04-01', 1, 0, ''),
(3, 'Ahmad Sudibyo', 'RJP88128739108', 'Serang', '1994-02-26', '2021-04-01', 2, 0, ''),
(7, 'Subekti', 'RJP321321', 'serang', '2021-04-05', '2021-04-21', 1, 1, 'asdas'),
(8, 'Nia Kurniawan', 'RJP12341234', 'Julikah', '2021-04-13', '2021-04-05', 1, 1, 'Lokino'),
(11, 'Basuki Kuswara', 'RJP123412342', 'Sidoarjo', '2021-04-20', '2021-04-28', 2, 7, 'Abasah');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `data_pegawai`
--
ALTER TABLE `data_pegawai`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `nip` (`nip`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `data_pegawai`
--
ALTER TABLE `data_pegawai`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
