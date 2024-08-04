<?php
session_start();

// 配置
define('ROOT_DIR', '/home/vol19_1/my-place.us/mp_37036943/htdocs');
ini_set('display_errors', 1);
error_reporting(E_ALL);

class FileManager {
    private $current_dir;
    private $message = '';

    public function __construct($dir = '') {
        $this->setCurrentDir($dir);
    }

    private function setCurrentDir($dir) {
        $path = realpath(ROOT_DIR . '/' . $dir);
        $this->current_dir = ($path && strpos($path, ROOT_DIR) === 0) ? $path : ROOT_DIR;
    }

    public function getCurrentDir() {
        return $this->current_dir;
    }

    public function getCurrentPath() {
        return str_replace(ROOT_DIR, '', $this->current_dir);
    }

    public function getMessage() {
        return $this->message;
    }

    public function getParentDir() {
        $parent = dirname($this->current_dir);
        return ($parent === $this->current_dir) ? '' : str_replace(ROOT_DIR, '', $parent);
    }

    public function getFileList() {
        $files = scandir($this->current_dir);
        $file_list = [];
        foreach ($files as $file) {
            if ($file != "." && $file != "..") {
                $full_path = $this->current_dir . '/' . $file;
                $file_list[] = [
                    'name' => $file,
                    'path' => str_replace(ROOT_DIR, '', $full_path),
                    'type' => is_dir($full_path) ? 'dir' : 'file',
                    'size' => is_dir($full_path) ? '-' : $this->humanFilesize(filesize($full_path)),
                    'permissions' => substr(sprintf('%o', fileperms($full_path)), -4)
                ];
            }
        }
        return $file_list;
    }

    private function humanFilesize($bytes, $decimals = 2) {
        $size = array('B','KB','MB','GB','TB','PB','EB','ZB','YB');
        $factor = floor((strlen($bytes) - 1) / 3);
        return sprintf("%.{$decimals}f", $bytes / pow(1024, $factor)) . @$size[$factor];
    }

    public function uploadFile($file) {
        $target_file = $this->current_dir . '/' . basename($file["name"]);
        if (move_uploaded_file($file["tmp_name"], $target_file)) {
            $this->message = "文件 ". basename($file["name"]). " 已上传成功。";
        } else {
            $this->message = "抱歉，文件上传失败。";
        }
    }

    public function createDirectory($name) {
        $path = $this->current_dir . '/' . $name;
        if (!file_exists($path)) {
            if (mkdir($path)) {
                $this->message = "目录创建成功。";
            } else {
                $this->message = "目录创建失败。";
            }
        } else {
            $this->message = "目录已存在。";
        }
    }

    public function renameFile($old_name, $new_name) {
        $old_path = $this->current_dir . '/' . $old_name;
        $new_path = $this->current_dir . '/' . $new_name;
        if (rename($old_path, $new_path)) {
            $this->message = "重命名成功。";
        } else {
            $this->message = "重命名失败。";
        }
    }

    public function deleteFile($name) {
        $path = $this->current_dir . '/' . $name;
        if (is_dir($path)) {
            if ($this->deleteDirectory($path)) {
                $this->message = "目录删除成功。";
            } else {
                $this->message = "目录删除失败。";
            }
        } else {
            if (unlink($path)) {
                $this->message = "文件删除成功。";
            } else {
                $this->message = "文件删除失败。";
            }
        }
    }

    private function deleteDirectory($dir) {
        if (!file_exists($dir)) return true;
        if (!is_dir($dir)) return unlink($dir);
        foreach (scandir($dir) as $item) {
            if ($item == '.' || $item == '..') continue;
            if (!$this->deleteDirectory($dir . DIRECTORY_SEPARATOR . $item)) return false;
        }
        return rmdir($dir);
    }

    public function copyFile($name) {
        $_SESSION['clipboard'] = [
            'action' => 'copy',
            'file' => $this->current_dir . '/' . $name
        ];
        $this->message = "文件已复制到剪贴板。";
    }

    public function cutFile($name) {
        $_SESSION['clipboard'] = [
            'action' => 'cut',
            'file' => $this->current_dir . '/' . $name
        ];
        $this->message = "文件已剪切到剪贴板。";
    }

    public function pasteFile() {
        if (isset($_SESSION['clipboard'])) {
            $clipboard = $_SESSION['clipboard'];
            $source = $clipboard['file'];
            $destination = $this->current_dir . '/' . basename($source);
            
            if ($clipboard['action'] == 'copy') {
                if (copy($source, $destination)) {
                    $this->message = "文件已成功复制。";
                } else {
                    $this->message = "复制文件失败。";
                }
            } elseif ($clipboard['action'] == 'cut') {
                if (rename($source, $destination)) {
                    $this->message = "文件已成功移动。";
                    unset($_SESSION['clipboard']);
                } else {
                    $this->message = "移动文件失败。";
                }
            }
        } else {
            $this->message = "剪贴板为空。";
        }
    }

    public function downloadFile($name) {
        $file_path = $this->current_dir . '/' . $name;
        if (file_exists($file_path)) {
            header('Content-Description: File Transfer');
            header('Content-Type: application/octet-stream');
            header('Content-Disposition: attachment; filename="'.basename($file_path).'"');
            header('Expires: 0');
            header('Cache-Control: must-revalidate');
            header('Pragma: public');
            header('Content-Length: ' . filesize($file_path));
            flush();
            readfile($file_path);
            exit;
        }
    }

    public function unzipFile($zipfile, $extractpath) {
        $zip = new ZipArchive;
        $zipfile_path = $this->current_dir . '/' . $zipfile;
        $extract_path = $this->current_dir . '/' . $extractpath;
        if ($zip->open($zipfile_path) === TRUE) {
            $zip->extractTo($extract_path);
            $zip->close();
            $this->message = "ZIP文件解压成功到：" . $extractpath;
        } else {
            $this->message = "无法打开ZIP文件。";
        }
    }
}

// 处理请求
$fm = new FileManager($_GET['dir'] ?? '');

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    if (isset($_FILES['fileToUpload'])) {
        $fm->uploadFile($_FILES['fileToUpload']);
    } elseif (isset($_POST['create_dir'])) {
        $fm->createDirectory($_POST['dir_name']);
    } elseif (isset($_POST['rename_file'])) {
        $fm->renameFile($_POST['old_name'], $_POST['new_name']);
    } elseif (isset($_POST['delete_file'])) {
        $fm->deleteFile($_POST['file_name']);
    } elseif (isset($_POST['copy_file'])) {
        $fm->copyFile($_POST['file_name']);
    } elseif (isset($_POST['cut_file'])) {
        $fm->cutFile($_POST['file_name']);
    } elseif (isset($_POST['paste_file'])) {
        $fm->pasteFile();
    } elseif (isset($_POST['unzip_file'])) {
        $fm->unzipFile($_POST['zipfile'], $_POST['extractpath']);
    }
}

if (isset($_GET['download'])) {
    $fm->downloadFile($_GET['download']);
}

$file_list = $fm->getFileList();
$current_dir = $fm->getCurrentDir();
$current_path = $fm->getCurrentPath();
$parent_dir = $fm->getParentDir();
$message = $fm->getMessage();
?>

<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>文件管理器</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            margin: 0;
            padding: 20px;
            background-color: #f4f4f4;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
            background-color: #fff;
            padding: 20px;
            border-radius: 5px;
            box-shadow: 0 0 10px rgba(0,0,0,0.1);
        }
        h1, h2 {
            color: #333;
        }
        table {
            border-collapse: collapse;
            width: 100%;
            margin-bottom: 20px;
        }
        th, td {
            border: 1px solid #ddd;
            padding: 12px;
            text-align: left;
        }
        th {
            background-color: #f2f2f2;
            font-weight: bold;
        }
        tr:nth-child(even) {
            background-color: #f9f9f9;
        }
        a {
            color: #1a73e8;
            text-decoration: none;
        }
        a:hover {
            text-decoration: underline;
        }
        .btn {
            display: inline-block;
            padding: 8px 12px;
            background-color: #1a73e8;
            color: white;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            font-size: 14px;
        }
        .btn:hover {
            background-color: #185abc;
        }
        input[type="text"], input[type="file"] {
            padding: 8px;
            margin-right: 10px;
            border: 1px solid #ddd;
            border-radius: 4px;
        }
        .message {
            background-color: #e7f3fe;
            border-left: 6px solid #2196F3;
            margin-bottom: 15px;
            padding: 10px;
            color: #0c5460;
        }
        .operations {
            display: flex;
            gap: 10px;
            align-items: center;
        }
        .operations form {
            margin: 0;
        }
        .breadcrumb {
            margin-bottom: 20px;
            padding: 10px;
            background-color: #f8f9fa;
            border-radius: 4px;
        }
        .breadcrumb a {
            color: #007bff;
            text-decoration: none;
        }
        .breadcrumb a:hover {
            text-decoration: underline;
        }
    </style>
    <script>
        function confirmDelete(fileName) {
            return confirm("确定要删除 " + fileName + " 吗？");
        }
    </script>
</head>
<body>
    <div class="container">
        <h1>文件管理器</h1>
        
        <?php if ($message): ?>
            <div class="message"><?php echo htmlspecialchars($message); ?></div>
        <?php endif; ?>

        <h2>当前目录: <?php echo htmlspecialchars($current_path ? $current_path : '/'); ?></h2>

        <div class="breadcrumb">
            <a href="?">根目录</a>
            <?php
            $path_parts = explode('/', trim($current_path, '/'));
            $cumulative_path = '';
            foreach ($path_parts as $part) {
                $cumulative_path .= '/' . $part;
                echo ' / <a href="?dir=' . urlencode(trim($cumulative_path, '/')) . '">' . htmlspecialchars($part) . '</a>';
            }
            ?>
        </div>

        <?php if ($parent_dir !== ''): ?>
            <p><a href="?dir=<?php echo urlencode($parent_dir); ?>" class="btn">返回上级目录</a></p>
        <?php endif; ?>

        <table>
            <tr>
                <th>名称</th>
                <th>类型</th>
                <th>大小</th>
                <th>权限</th>
                <th>操作</th>
            </tr>
            <?php foreach ($file_list as $file): ?>
                <tr>
                    <td>
                        <?php if ($file['type'] == 'dir'): ?>
                            <a href="?dir=<?php echo urlencode($file['path']); ?>"><?php echo htmlspecialchars($file['name']); ?></a>
                        <?php else: ?>
                            <?php echo htmlspecialchars($file['name']); ?>
                        <?php endif; ?>
                    </td>
                    <td><?php echo htmlspecialchars($file['type']); ?></td>
                    <td><?php echo htmlspecialchars($file['size']); ?></td>
                    <td><?php echo htmlspecialchars($file['permissions']); ?></td>
                    <td class="operations">
                        <?php if ($file['type'] != 'dir'): ?>
                            <a href="?download=<?php echo urlencode($file['name']); ?>" class="btn">下载</a>
                        <?php endif; ?>
                        <form method="post" style="display:inline;">
                            <input type="hidden" name="old_name" value="<?php echo htmlspecialchars($file['name']); ?>">
                            <input type="text" name="new_name" placeholder="新名称">
                            <input type="submit" name="rename_file" value="重命名" class="btn">
                        </form>
                        <form method="post" style="display:inline;" onsubmit="return confirmDelete('<?php echo htmlspecialchars($file['name']); ?>')">
                            <input type="hidden" name="file_name" value="<?php echo htmlspecialchars($file['name']); ?>">
                            <input type="submit" name="delete_file" value="删除" class="btn">
                        </form>
                        <form method="post" style="display:inline;">
                                                        <input type="hidden" name="file_name" value="<?php echo htmlspecialchars($file['name']); ?>">
                            <input type="submit" name="copy_file" value="复制" class="btn">
                            <input type="submit" name="cut_file" value="剪切" class="btn">
                        </form>
                    </td>
                </tr>
            <?php endforeach; ?>
        </table>

        <form method="post" style="margin-bottom: 20px;">
            <input type="submit" name="paste_file" value="粘贴" class="btn">
        </form>

        <h2>上传文件</h2>
        <form action="" method="post" enctype="multipart/form-data">
            <input type="file" name="fileToUpload" id="fileToUpload">
            <input type="submit" value="上传文件" name="submit" class="btn">
        </form>

        <h2>创建目录</h2>
        <form method="post">
            <input type="text" name="dir_name" placeholder="新目录名">
            <input type="submit" name="create_dir" value="创建目录" class="btn">
        </form>

        <h2>ZIP解压工具</h2>
        <form method="post">
            <input type="text" name="zipfile" placeholder="输入ZIP文件名">
            <input type="text" name="extractpath" placeholder="输入解压目标路径">
            <input type="submit" name="unzip_file" value="解压文件" class="btn">
        </form>
    </div>
</body>
</html>