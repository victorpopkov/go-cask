cask 'example-two' do
  if MacOS.release <= :el_capitan
    version '1.5.0'
    sha256 '1f4dc096d58f7d21e3875671aee6f29b120ab84218fa47db2cb53bc9eb5b4dac'

    url "https://example.com/app_#{version}.pkg"
    appcast "https://example.com/sparkle/#{version}/el_capitan.xml",
            checkpoint: '93ef3101ca730028d70524f71b7f6f17cbdb8d26906299f90c38b7079e1d03a4'
  else
    version '2.0.0'
    sha256 'f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261'

    url "https://example.com/app_#{version}.pkg"
    appcast "https://example.com/sparkle/#{version.major}/appcast.xml",
            checkpoint: '57956bd3fb23a5673e30dc83ed19d51b43e5a9235756887f3ed90662e6c68fb7'
  end

  name 'Example'
  name 'Example Two'
  homepage 'https://example.com/'

  pkg "app_#{version}.pkg", allow_untrusted: true

  uninstall pkgutil: 'com.example.pkg.*'
end
