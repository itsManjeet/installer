#include "bootloader.hh"

#include <filesystem>

#include "../../../disk/disk.hh"
#include "../../../utils/exec.hh"

bool Bootloader::process() {
  if (m_Data->isEfi()) {
    std::string efiDir = m_Data->workDir() + "/boot/efi";
    std::error_code err;

    std::filesystem::create_directory(efiDir, err);
    if (err) {
      m_Mesg = "Failed to prepare EFI dir, " + err.message();
      return false;
    }

    auto [status, output] =
        Exec::output(("mount " + m_Data->bootDevice() + " " + efiDir).c_str());
    if (status != 0) {
      m_Mesg = "Failed to prepare EFI mount, " + output;
      return false;
    }
  }

  auto [status, output] = Exec::output(
      ("grub-install --root-directory=" + m_Data->workDir() +
       " --boot-directory=" + m_Data->workDir() + "/boot --recheck " +
       (m_Data->isEfi() ? "" : m_Data->bootDevice()))
          .c_str());
  if (status != 0) {
    m_Mesg = "Failed to install bootloader, " + output;
    return false;
  }

  std::string const GRUB_CONIG = R"(
insmod part_gpt
insmod part_msdos
insmod all_video
timeout=5
default='rlxos [inital-setup]'

menuentry 'rlxos [inital-setup]' {
  insmod gzio
  insmod ext2
  linux /boot/vmlinuz-%s root=UUID=%s system=%s
  initrd /boot/initrd-%s
})";

  std::string kernel_version;
  {
    auto [status, output] = Exec::output("uname -r");
    if (status != 0) {
      m_Mesg = "Failed to get kernel version," + output;
      return false;
    }

    kernel_version = output;
  }

  auto grub_cfg =
      fopen((m_Data->workDir() + "/boot/grub/grub.cfg").c_str(), "w");

  if (grub_cfg == nullptr) {
    m_Mesg =
        "Failed to write grub configuration, " + std::string(strerror(errno));
    return false;
  }

  auto disk = Disk(m_Data->disk());

  fprintf(grub_cfg, GRUB_CONIG.c_str(), kernel_version.c_str(),
          disk.uuid().c_str(), m_Data->version().c_str(),
          kernel_version.c_str());
  fclose(grub_cfg);

  m_Mesg = "configured bootloader";
  return true;
}