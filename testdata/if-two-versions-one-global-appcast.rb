cask 'if-two-versions-one-global-appcast' do
  if MacOS.release == :mavericks
    version '1.0.0'
    sha256 '92521fc3cbd964bdc9f584a991b89fddaa5754ed1cc96d6d42445338669c1305'
  else
    version '2.0.0'
    sha256 'f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261'
  end

  url "https://example.com/app_#{version}.dmg"
  appcast "https://example.com/sparkle/#{version.major}/appcast.xml",
          checkpoint: '8dc47a4bcec6e46b79fb6fc7b84224f1461f18a2d9f2e5adc94612bb9b97072d'
  name 'Example'
  name 'Example (if-two-versions-one-global-appcast)'
  homepage 'https://example.com/'
  license :commercial

  auto_updates true

  app 'Example (if-two-versions-one-global-appcast).app', target: 'Example.app'
  binary "#{appdir}/Example.app/Contents/MacOS/example-if", target: 'example'
end
