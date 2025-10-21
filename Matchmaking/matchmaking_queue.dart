import 'dart:async';
import 'dart:io';
import 'dart:math';

class Generator {
  static const _chars =
      'AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz1234567890';
  static Random _rnd = Random();

  static String id() => String.fromCharCodes(Iterable.generate(
      12, (_) => _chars.codeUnitAt(_rnd.nextInt(_chars.length))));

  static int level() => _rnd.nextInt(25);
}

class Player {
  final String id;
  final int level;
  const Player({required this.id, required this.level});

  @override
  String toString() {
    return 'Player(id: $id,\n\tlevel:$level);';
  }
}

class Match {
  final int refLevel; // metadaat.
  final Set<Player> players;
  const Match({required this.refLevel, required this.players});
}

abstract class PlayerManager {
  Stream<Player> getPlayer();
}

class FakePlayerManager implements PlayerManager {
  final _stream = Stream.periodic(const Duration(seconds: 1),
      (n) => Player(id: Generator.id(), level: Generator.level()));
  @override
  Stream<Player> getPlayer() {
    return _stream;
  }
}

class MatchmakingQueue {
  final PlayerManager playerManager;
  final Map<String, Player> _searchingPlayers = {};

  StreamSubscription? _playerSubscription;
  Timer? _findMatchTimer;

  Future<void> playerAdded(Player player) async {
    print('Adding player: $player.');
    print('Now there are a total of: ${_searchingPlayers.length + 1}');
    _searchingPlayers[player.id] = player;
  }

  bool isEligible(Player player, int refLevel) {
    return (player.level - refLevel).abs() < 6;
  }

  Future<Match?> findMatch() async {
    if (_searchingPlayers.isEmpty || _searchingPlayers.entries.length < 6) {
      return null;
    }
    final refLevel = _searchingPlayers.values.first.level;
    final potential = <Player>{};
    for (final player in _searchingPlayers.values) {
      if (isEligible(player, refLevel)) {
        potential.add(player);
      }
      if (potential.length == 6) {
        for (final p in potential) {
          _searchingPlayers.remove(p.id);
        }
        return Match(refLevel: refLevel, players: potential);
      }
    }
    return null;
  }

  MatchmakingQueue({required this.playerManager}) {
    _playerSubscription = playerManager.getPlayer().listen(playerAdded);
    _findMatchTimer =
        Timer.periodic(const Duration(seconds: 1), (t) => findMatch());
  }

  void close() {
    _findMatchTimer?.cancel();
    _playerSubscription?.cancel();
  }
}

Future<void> queueLoop(MatchmakingQueue queue) async {
  while (true) {
    print('Searching for match...');
    final match = await queue.findMatch();
    print('Match found: $match');
    await Future.delayed(const Duration(seconds: 2));
  }
}

void main() async {
  final queue = MatchmakingQueue(playerManager: FakePlayerManager());

  await queueLoop(queue);
}
