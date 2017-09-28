cask 'example-one' do
  version '2.0.0'
  sha256 'f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261'

  url "https://example.com/app_#{version}.dmg"
  appcast "https://example.com/sparkle/#{version.major}/appcast.xml",
          checkpoint: '8dc47a4bcec6e46b79fb6fc7b84224f1461f18a2d9f2e5adc94612bb9b97072d'
  name 'Example'
  name 'Example One'
  homepage 'https://example.com/'

  auto_updates true

  app "Example #{version.major_minor}.app", target: 'Example.app'
  app "Example #{version.major_minor} Uninstaller.app"
  binary "#{appdir}/Example #{version.major_minor}.app/Contents/MacOS/example-one", target: 'example'
end
