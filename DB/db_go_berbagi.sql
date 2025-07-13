-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jul 13, 2025 at 06:47 AM
-- Server version: 8.0.30
-- PHP Version: 8.2.17

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_go_berbagi`
--

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `occupation` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `avatar_file_name` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password_hash`, `avatar_file_name`, `role`, `created_at`, `updated_at`) VALUES
(1, 'Awan', 'Programmer', 'awan@g.com', '123123', 'http://www.google.com', 'admin', '2025-07-08 16:08:08', '2025-07-08 16:08:08'),
(2, 'test simpan', '', '', '', '', '', '2025-07-08 16:46:34', '2025-07-08 16:46:34'),
(3, 'Awan', 'Waiters', 'awan@g.com', '$2a$04$HnkCCMuGGCPIpekiK3hFnOaJ1sk5NP8kSFGqw60iS9/zMhf/UUwcy', '', 'user', '2025-07-09 16:48:29', '2025-07-09 16:48:29'),
(4, 'Posman', 'masak', 'pos@g.com', '$2a$04$zW4DFEUIRO/cjUMgrL.Lo.3ieDWNqCI3znu8nzfCPofBTmjOo8DMC', '', 'user', '2025-07-09 17:07:37', '2025-07-09 17:07:37'),
(10, 'Posman2', 'masak', 'pos2@g.com', '$2a$04$5srNPnYCmKhJZxiYu9GYseIRAsc6D0kFquoRQxU3IUdLIUJmzRqSK', '', 'user', '2025-07-10 15:45:23', '2025-07-10 15:45:23'),
(11, 'Posman2', 'masak', 'pos2@g.com', '$2a$04$L7B3C6/ul6beUVK6vFQtL.ayDj7d7j1PAsqkoliMNoClKvJJx6WKi', '', 'user', '2025-07-10 15:48:18', '2025-07-10 15:48:18'),
(12, 'Posman3', 'masak', 'pos3@g.com', '$2a$04$PsQzJrzKFtfJTQCNuafXVehpJqYO75tNQ.sR9mcasJiEDbs.XXLYS', '', 'user', '2025-07-10 15:51:37', '2025-07-10 15:51:37'),
(13, 'Posman4', 'masak', 'pos4@g.com', '$2a$04$C/qUyKB6S/bXquK5BIUbDODo0Lc/5dfPw8yAF1NWXPDQeNUnEVIQm', '', 'user', '2025-07-10 16:36:24', '2025-07-10 16:36:24'),
(14, 'Posman4', 'masak', 'pos4@g.com', '$2a$04$7wi1T12BdYo4u.tWZaOSL.41KDLU0Va3FCpcNvIckVM4FblSIfSni', '', 'user', '2025-07-10 16:36:57', '2025-07-10 16:36:57');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
