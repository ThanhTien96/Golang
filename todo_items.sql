/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE TABLE `todo_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `description` text,
  `status` varchar(50) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `todo_items` (`id`, `title`, `description`, `status`, `created_at`, `updated_at`) VALUES
(1, 'todo 01', 'description for todo list 01', 'Doing', '2024-01-26 04:53:22', '2024-01-26 09:27:48');
INSERT INTO `todo_items` (`id`, `title`, `description`, `status`, `created_at`, `updated_at`) VALUES
(2, 'todo 02', 'description for todo list 02', 'Doing', '2024-01-26 04:54:47', '2024-01-26 10:49:26');
INSERT INTO `todo_items` (`id`, `title`, `description`, `status`, `created_at`, `updated_at`) VALUES
(4, 'todo 03', 'description for todo list 03', 'Doing', '2024-01-26 07:39:39', '2024-01-26 09:27:48');
INSERT INTO `todo_items` (`id`, `title`, `description`, `status`, `created_at`, `updated_at`) VALUES
(5, 'todo 04', 'description for todo list 04', 'Doing', '2024-01-26 07:39:44', '2024-01-26 09:27:48'),
(6, 'todo 05', 'description for todo list 05', 'Doing', '2024-01-26 07:39:50', '2024-01-26 09:27:48'),
(7, 'todo 06', 'description for todo list 06', 'Doing', '2024-01-26 07:40:00', '2024-01-26 09:27:48'),
(8, 'todo 07', 'description for todo list 07', 'Doing', '2024-01-26 07:40:06', '2024-01-26 09:27:48'),
(9, 'todo 08', 'description for todo list 08', 'Doing', '2024-01-26 10:13:02', '2024-01-26 10:49:19'),
(10, 'todo 09', 'description for todo list 09', 'Done', '2024-01-26 10:13:20', NULL);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;