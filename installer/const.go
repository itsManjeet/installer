package installer

var REQUIRED_TOOLS = []string{"unsquashfs", "grub-install", "mkdir", "mount", "useradd", "openssl"}

var REQUIRED_DIR = []string{"work", "cache", "root", "sysroot", "boot", "apps", "home"}

var GRUB_CONFIG = `
insmod part_gpt
insmod part_msdos
insmod all_video
timeout=5
default='rlxos [inital-setup]'
menuentry 'rlxos [inital-setup]' {
  insmod gzio
  insmod ext2
  linux /boot/vmlinuz-%s root=UUID=%s quiet splash fastboot
  initrd /boot/initrd-%s
})
`
