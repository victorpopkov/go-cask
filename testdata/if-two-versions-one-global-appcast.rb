cask 'if-two-versions-one-global-appcast' do
  if MacOS.version == :mavericks
    version '1.0.0'
    sha256 '92521fc3cbd964bdc9f584a991b89fddaa5754ed1cc96d6d42445338669c1305'
  else
    version '2.0.0'
    sha256 'f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261'
  end

  url "https://example.com/app_#{version}.dmg"
  appcast "https://example.com/sparkle/#{version.major}/appcast.xml"
  name 'Example'
  name 'Example (if-two-versions-one-global-appcast)'
  homepage 'https://example.com/'

  auto_updates true

  app 'Example (if-two-versions-one-global-appcast).app', target: 'Example.app'
  binary "#{appdir}/Example.app/Contents/MacOS/example-if", target: 'example'
end
